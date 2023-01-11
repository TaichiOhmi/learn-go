package app2

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	_, err := IsEmail("hello")
	if err == nil {
		t.Error("hello is not an email")
	}
	_, err = IsEmail("taichi@gml.com")
	if err == nil {
		t.Error("taichi@gml.com is an email")
	}
	_, err = IsEmail("taichi@gml")
	if err != nil {
		t.Error("taichi@gml is an email")
	}
}
