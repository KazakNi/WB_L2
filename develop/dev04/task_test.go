package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnagram(t *testing.T) {
	type test struct {
		input  []string
		output map[string][]string
	}

	tests := []test{
		{input: []string{"пятак", "пятка", "пятка", "тяпка", "листок", "слиток", "столик"},
			output: map[string][]string{
				"листок": {"листок", "слиток", "столик"},
				"пятак":  {"пятак", "пятка", "тяпка"},
			},
		},
		{input: []string{"Ласков", "словак", "славок", "сковал", "Марина", "Армани", "ранами", "ранима"},
			output: map[string][]string{
				"ласков": {"ласков", "словак", "славок", "сковал"},
				"марина": {"марина", "армани", "ранами", "ранима"},
			},
		},
	}

	for _, tc := range tests {
		got := AnagramSet(tc.input)
		assert.EqualValues(t, tc.output, *got)
	}
}
