package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunMain(t *testing.T) {
	main()
	assert.True(t, true, "This is good. Canary test.")
}
