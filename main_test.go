package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "odd length palindrome",
			input:    "racecar",
			expected: true,
		},
		{
			name:     "even length palindrome",
			input:    "abba",
			expected: true,
		},
		{
			name:     "not a palindrome",
			input:    "hello",
			expected: false,
		},
		{
			name:     "single character",
			input:    "a",
			expected: true,
		},
		{
			name:     "empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "two different characters",
			input:    "ab",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsPalindrome(test.input)

			if result != test.expected {
				t.Errorf(
					"IsPalindrome(%q) = %t, want %t",
					test.input,
					result,
					test.expected,
				)
			}
		})
	}
}
