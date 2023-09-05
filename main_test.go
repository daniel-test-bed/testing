package main

import (
	"github.com/stretchr/testify/assert"
	tst "testing"
)

func TestRunMain(t *tst.T) {
	main()
	assert.True(t, true, "This is good. Canary test.")
}
