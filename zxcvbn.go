package zxcvbn

import (
	"time"

	"go.lucor.dev/zxcvbn/match"
	"go.lucor.dev/zxcvbn/matching"
	"go.lucor.dev/zxcvbn/scoring"
	"go.lucor.dev/zxcvbn/utils/math"
)

// PasswordStrength takes a password, userInputs and optional filters and returns a MinEntropyMatch
func PasswordStrength(password string, userInputs []string, filters ...func(match.Matcher) bool) scoring.MinEntropyMatch {
	start := time.Now()
	matches := matching.Omnimatch(password, userInputs, filters...)
	result := scoring.MinimumEntropyMatchSequence(password, matches)
	end := time.Now()

	calcTime := end.Nanosecond() - start.Nanosecond()
	result.CalcTime = zxcvbnmath.Round(float64(calcTime)*time.Nanosecond.Seconds(), .5, 3)
	return result
}
