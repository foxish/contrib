package comment

import "github.com/google/go-github/github"

func FindComments(events []*github.IssueComment, matcher Matcher) []*github.IssueComment {
	matchingComments := []*github.IssueComment{}

	for _, event := range events {
		if matcher.Match(event) {
			matchingComments = append(matchingComments, event)
		}
	}

	return matchingComments
}
