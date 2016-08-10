package comment

import "github.com/google/go-github/github"

func FindComment(events []*github.Comment, matcher Matcher) []*github.Comment {
	matchingComments := []*github.Comment{}

	for _, event := range events {
		if matcher.Match(event) {
			matchingComments = append(matchingComments, event)
		}
	}

	return matchingComments
}
