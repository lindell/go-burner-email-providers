package burner_test

import (
	"testing"

	"github.com/lindell/go-burner-email-providers/burner"
)

func TestNonBurnerEmails(t *testing.T) {
	nonBurnerEmails := []string{
		"hello@gmail.com",
		"my.name@hotmail.com",
		"hello@outlook.com",
		"hello@protonmail.com",
	}
	for _, email := range nonBurnerEmails {
		if ok := burner.IsBurnerEmail(email); ok {
			t.Errorf("%s should not be considered burner email", email)
		}
	}
}

func TestBurnerEmails(t *testing.T) {
	burnerEmails := []string{
		"name@burnermail.io",
		"yo@yopmail.com",
		"abcd.com",
	}
	for _, email := range burnerEmails {
		if ok := burner.IsBurnerEmail(email); !ok {
			t.Errorf("%s should be considered burner email", email)
		}
	}
}

func TestNonBurnerDomains(t *testing.T) {
	nonBurnerDomains := []string{
		"gmail.com",
		"hotmail.com",
		"outlook.com",
		"protonmail.com",
	}
	for _, email := range nonBurnerDomains {
		if ok := burner.IsBurnerDomain(email); ok {
			t.Errorf("%s should not be considered burner domain", email)
		}
	}
}

func TestBurnerDomains(t *testing.T) {
	burnerEmails := []string{
		"burnermail.io",
		"yopmail.com",
	}
	for _, email := range burnerEmails {
		if ok := burner.IsBurnerDomain(email); !ok {
			t.Errorf("%s should be considered burner domain", email)
		}
	}
}
