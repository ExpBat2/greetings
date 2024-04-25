package greetings

import (
	"regexp"
	"testing"
)

func testHello(t *testing.T) {
	name := "Javier"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Javier")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Juan") = %q,%v, quiere coincidencia para %#q,nil`, msg, err, want)
	}
}

func testHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("")= %q,%v,quiere un "", error`, msg, err)
	}
}
