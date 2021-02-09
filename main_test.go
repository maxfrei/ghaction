package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	expected := 12
	result := sum(5, 7)
	assert.Equal(t, result, expected)
}
