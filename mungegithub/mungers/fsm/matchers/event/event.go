package main

import "github.com/google/go-github/github"

type Matcher interface {
	Match(event *github.Event) bool
}

type And []Matcher

func (a And) Match(event *github.Event) bool {
	for _, matcher := range []Matcher(a) {
		if !matcher.Match(event) {
			return false
		}
	}
	return true
}

type Or []Matcher

func (o Or) Match(event *github.Event) bool {
	for _, matcher := range []Matcher(o) {
		if matcher.Match(event) {
			return true
		}
	}
	return false
}

type Not struct {
	Matcher Matcher
}

func (n Not) Match(event *github.Event) bool {
	return !n.Matcher.Match(event)
}
