package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSorting(t *testing.T) {
	type test struct {
		params Params
		file   string
		res    [][]string
		err    error
	}

	tests := []test{
		{file: "test.txt", res: [][]string{
			{"ahahaha", "crazy", "test"},
			{"bimbaza", "bombata"},
			{"lalalala", "alas"},
			{"lalalala", "alas"},
			{"lol", "kek", "cheburek"},
		},
			err: nil, params: Params{}},
		{file: "test.txt", res: [][]string{
			{"lol", "kek", "cheburek"},
			{"lalalala", "alas"},
			{"lalalala", "alas"},
			{"bimbaza", "bombata"},
			{"ahahaha", "crazy", "test"},
		},
			err: nil, params: Params{r: true}},
		{file: "test.txt", res: [][]string{
			{"lalalala", "alas"},
			{"bimbaza", "bombata"},
			{"ahahaha", "crazy", "test"},
			{"lol", "kek", "cheburek"},
		},
			err: nil, params: Params{u: true, k: 1}},
	}

	for _, tc := range tests {
		got := SortMan(tc.file, tc.params)
		assert.ElementsMatch(t, got, tc.res)
	}
}
