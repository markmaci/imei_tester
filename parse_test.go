package main

import (
	"errors"
	"testing"
)

func TestValidate(t *testing.T) {
	got, err := validateChecksum("868880045952250")
	want, expected := false, errors.New("error = invalid checksum")

	if want != got {
		t.Errorf("Expected '%t', but got '%t", want, got)
	}

	if err != expected {
		t.Log(err)
	}
}

func TestValidate2(t *testing.T) {
	got, err := validateChecksum("357897085527974")
	want := true

	if want != got {
		t.Errorf("Expected '%t', but got '%t", want, got)
		t.Log(err)
	}
}

func TestParse(t *testing.T) {
	got, err := parseIMEI("357700101512527")
	want := parsedIMEI{typeAllocationCode: "35770010", serialNumber: "151252", checksum: 7}

	if got != want {
		t.Errorf("failed")
		t.Log(err)
	}
}

func TestParse2(t *testing.T) {
	got, err := parseIMEI("861663039053692")
	want := parsedIMEI{typeAllocationCode: "86166303", serialNumber: "905369", checksum: 2}

	if got != want {
		t.Errorf("failed")
		t.Log(err)
	}
}
