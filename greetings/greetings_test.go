package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloLowerCase(t *testing.T) {
	name := "lowercase"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello("lowercase")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("lowercase") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloUpperCase(t *testing.T) {
	name := "UPPERCASE"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello("UPPERCASE")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("UPPERCASE") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloMixedCase(t *testing.T) {
	name := "mIXedCase"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello("mIXedCase")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("mIXedCase") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloSpecialChars(t *testing.T) {
	name := "!@£$/|"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello("!@£$/|")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("!@£$/|") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloMixedString(t *testing.T) {
	name := "Boris! the 2()nd"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello("Boris! the 2nd)")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Boris! the 2nd)") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
