package event

import "github.com/google/go-github/github"

func FindEvent(events []*github.Event, matcher Matcher) []*github.Event {
	matchingEvents := []*github.Event{}

	for _, event := range events {
		if matcher.Match(event) {
			matchingEvents = append(matchingEvents, event)
		}
	}

	return matchingEvents
}
