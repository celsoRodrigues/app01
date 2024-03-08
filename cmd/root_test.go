package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestExecute(t *testing.T) {
	var buff bytes.Buffer
	rootCmd.SetOut(&buff)
	rootCmd.Execute()
	expected := "This application is a tool to generate the needed files"
	if !strings.Contains(strings.TrimSpace(buff.String()), expected) {
		t.Fatalf("got %v, expected %v", strings.TrimSpace(buff.String()), expected)
	}
}

func TestAddVersion(t *testing.T) {
	expected := true
	got := AddVersion("")
	if got != expected {
		t.Fatalf("expected %v got %v", expected, got)
	}
}
