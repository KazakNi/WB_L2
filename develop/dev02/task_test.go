package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnpack(t *testing.T) {
	type test struct {
		input string
		res   string
		err   error
	}

	tests := []test{
		{input: "a4bc2d5e", res: "aaaabccddddde", err: nil},
		{input: "abcd", res: "abcd", err: nil},
		{input: "45", res: "", err: ErrInvalidString},
	}

	for _, tc := range tests {
		got, err := UnpackString(tc.input)
		assert.Equal(t, got, tc.res)
		assert.Equal(t, err, tc.err)
	}
}
