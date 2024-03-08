package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestVersion(t *testing.T) {

	ttests := []struct {
		name     string
		expected string
		input    string
		command  string
	}{
		{
			name:     "Testing version command when value is passed",
			expected: "app01 version 0.1.000",
			input:    "0.1",
			command:  "version",
		},
		{
			name:     "Testing version command when value is empty",
			expected: "app01 version 000",
			input:    "",
			command:  "version",
		},
	}

	for _, test := range ttests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			var buff bytes.Buffer
			rootCmd.SetOut(&buff)
			AddVersion(test.input)
			rootCmd.SetArgs([]string{test.command})
			rootCmd.Execute()

			if strings.TrimSpace(buff.String()) != test.expected {
				t.Fatalf("expected: \"%v\" got: \"%v\"", test.expected, strings.TrimSpace(buff.String()))
			}
		})
	}
}
