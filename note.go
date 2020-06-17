package main

import "strings"

type note struct {
	content string
}

func (n *note) title() string {
	if n.content == "" {
		return "Untitled"
	}

	return strings.SplitN(n.content, "\n", 2)[0]
}