package comment

import (
	"fmt"
	"regexp"

	"github.com/google/go-github/github"
)

type BotMessage struct {
	Name string
}

func (b BotMessage) Match(comment *github.IssueComment) bool {
	if comment.Body == nil {
		return false
	}
	match, _ := regexp.MatchString(fmt.Sprintf("^[%s]", b.Name), *comment.Body)
	return match
}

type Command struct {
	Name string
}

func (c Command) Match(comment *github.IssueComment) bool {
	if comment.Body == nil {
		return false
	}
	match, _ := regexp.MatchString(fmt.Sprintf("^/%s", c.Name), *comment.Body)
	return match
}
