// Package burner provides functions to check if the domain of an email or a domain is in the community maintained https://github.com/wesbos/burner-email-providers burner list.
package burner

import (
	"strings"
)

// IsBurnerEmail checks if the domain of an email address is in a list of burner email domain names
func IsBurnerEmail(email string) bool {
	at := strings.LastIndex(email, "@")
	if at == -1 {
		return true // not a valid email,
	}
	domain := email[at+1:]

	_, ok := domains[domain]
	return ok
}

// IsBurnerDomain checks if a domain is in a list of burner email domain names
func IsBurnerDomain(domain string) bool {
	_, ok := domains[domain]
	return ok
}
