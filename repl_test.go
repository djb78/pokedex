package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input	string
		expected []string
	}{
		{
			input: "  hello world  ",
			expected: []string{"hello", "world"},
		},
	}
	
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("%d expected, got %d", len(c.expected), len(actual))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Fail()
				t.Errorf("expected %s, got %s", expectedWord, word) }
		}
	}
}
