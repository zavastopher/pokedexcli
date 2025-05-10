package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "word",
			expected: []string{"word"},
		},
		{
			input:    " word1     word2",
			expected: []string{"word1", "word2"},
		},
		{
			input:    "Pikachu Bulbasaur     Charmander Squirtle    ",
			expected: []string{"Pikachu", "Bulbasaur", "Charmander", "Squirtle"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the
		// expected slice, if they don't match, use t.Errorf to
		// print an error messag and fail the test.
		actualLen := len(actual)
		expectedLen := len(c.expected)
		if actualLen != expectedLen {
			t.Errorf(`The length of the cleaned input slice for input %v is %v which does not match expected %v length`, c.input, actualLen, expectedLen)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// check each word in the slice to make sure they match
			if word != expectedWord {
				t.Errorf(`The words in the cleaned input do not match\nExpected: %v\nOutput: %v`, expectedWord, word)
			}
		}
	}
}

func TestHelp(t *testing.T) {
	if error := commandHelp(nil); error != nil {
		t.Errorf("Commandhelp function is unable to print text")
	}
}

func TestExit(t *testing.T) {
	if os.Getenv("exit") == "1" {
		commandExit(nil)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestExit")
	cmd.Env = append(os.Environ(), "exit=1")

	if err := cmd.Run(); err != nil {
		t.Errorf("Did not exit gracefully")
	}

}
