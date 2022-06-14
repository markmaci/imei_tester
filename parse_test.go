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
	got, err2 := parseIMEI("357700101512527")
	want := parsedIMEI{typeAllocationCode: "35770010", serialNumber: "151252", checksum: 7}

	if got != want {
		t.Errorf("failed")
		t.Logf(err2.Error())
	}
}
