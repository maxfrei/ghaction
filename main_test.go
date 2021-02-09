package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	expected := 12
	result := sum(5, 7)
	assert.Equal(t, result, expected)
	fmt.Println("---- env:", os.Getenv("MAXFREI_TEST"))
	fmt.Println("---- repo:", os.Getenv("MAXFREI_TEST_REPO"))
}
