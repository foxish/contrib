package comment

import (
	"time"

	"github.com/google/go-github/github"
)

type Matcher interface {
	Match(comment *github.IssueComment) bool
}

type And []Matcher

func (a And) Match(comment *github.IssueComment) bool {
	for _, matcher := range []Matcher(a) {
		if !matcher.Match(comment) {
			return false
		}
	}
	return true
}

type Or []Matcher

func (o Or) Match(comment *github.IssueComment) bool {
	for _, matcher := range []Matcher(o) {
		if matcher.Match(comment) {
			return true
		}
	}
	return false
}

type Not struct {
	Matcher Matcher
}

func (n Not) Match(comment *github.IssueComment) bool {
	return !n.Matcher.Match(comment)
}

type CreatedAfter struct {
	Date *time.Time
}

func (c CreatedAfter) Match(comment *github.IssueComment) bool {
	return comment.CreatedAt.After(*c.Date)
}
