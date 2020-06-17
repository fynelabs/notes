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

type notelist struct {
	notes []*note
}

func (l *notelist) add() *note {
	n := &note{}
	l.notes = append([]*note{n}, l.notes...)
	return n
}

func (l *notelist) remove(n *note) {
	for i, note := range l.notes {
		if n == note {
			if i == len(l.notes) - 1 {
				l.notes = l.notes[:i]
				return
			}
			l.notes = append(l.notes[:i], l.notes[i+1:]...)
			return
		}
	}
}