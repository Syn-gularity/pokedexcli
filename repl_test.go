package main

import(
	"testing"
)



func TestCleanInput(t *testing.T) { // TestXXX with t *testing.T as argument
	cases := []struct {
		input 	string
		expected []string
	}{
		{
			input: "  hello  world  ",
			expected: []string{"hello","world"},
		},
		//more cases
		{
			input: "nothing special here!",
			expected: []string{"nothing","special","here!"},
		},
		{
			input: "BIG WORDS",
			expected: []string{"big","words"},
		},
		{
			input: "What about Numb3rs?",
			expected: []string{"what","about", "numb3rs?"},
		},
		{
			input: "",
			expected: []string{},
		},
		{
			input: "  ",
			expected: []string{},
		},
	}

	for _, c := range cases{
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf("inputlength unequal expected length")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord{
				t.Errorf("%s unequal %s", word, expectedWord)
			}
		}
	}
}