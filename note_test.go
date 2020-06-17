package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoteTitle(t *testing.T) {
	n := &note{"Some content"}
	assert.Equal(t, "Some content", n.title())

	n = &note{"line1\nline2"}
	assert.Equal(t, "line1", n.title())
}
