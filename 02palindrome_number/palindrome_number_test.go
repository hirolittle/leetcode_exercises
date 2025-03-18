package palindrome_number

import (
	"testing"
)

func TestPalindromeNumber(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{"PalindromePositive", 121, true},
		{"NonPalindromePositive", 123, false},
		{"SingleDigit", 7, true},
		{"NegativeNumber", -121, false},
		{"Zero", 0, true},
		{"TrailingZeroNonPalindrome", 10, false},
		{"LargePalindrome", 1234321, true},
		{"LargeNonPalindrome", 1234567, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.input)
			if got != tt.want {
				t.Errorf("isPalindrome(%d) = %t, want %t", tt.input, got, tt.want)
			}
		})
	}
}
