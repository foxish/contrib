package ui

import (
	"time"

	"github.com/google/go-github/github"

	mgh "k8s.io/contrib/mungegithub/github"
	"k8s.io/contrib/mungegithub/mungers/fsm/matchers/comment"
)

type Pinger struct {
	keyword     string        // Short description for the ping
	description string        // Long description for the ping
	startDate   *time.Time    // Pinger started at this time
	timePeriod  time.Duration // How often should we ping
	maxCount    int           // Will stop pinging after that many times
}

func newPinger(keyword string) *Pinger {
	return &Pinger{
		keyword: keyword,
	}
}

// SetDescription is the description that goes along the ping
func (p *Pinger) SetDescription(description string) *Pinger {
	p.description = description

	return p
}

// SetTimePeriod is the time we wait between pings
func (p *Pinger) SetTimePeriod(timePeriod time.Duration) *Pinger {
	p.timePeriod = timePeriod

	return p
}

// SetMaxCount will make the pinger fail when it reaches maximum
func (p *Pinger) SetMaxCount(maxCount int) *Pinger {
	p.maxCount = maxCount

	return p
}

func (p *Pinger) SetStartDate(date *time.Time) *Pinger {
	p.startDate = date

	return p
}

// Ping who on the object! Returns true if we reached the ping limit
func (p *Pinger) Ping(obj *mgh.MungeObject, who string) (bool, error) {
	comments, err := obj.ListComments()
	if err != nil {
		return false, err
	}

	pings := comment.FindComments(comments, p.getMatcher())

	// We have pinged too many times, it's time to try something else
	if p.maxReached(pings) {
		return true, nil
	}

	if !p.shouldPingNow(pings) {
		return false, nil
	}

	return false, Comment(obj, p.keyword, who, p.description)
}

func (p *Pinger) getMatcher() comment.Matcher {
	if p.startDate != nil {
		return comment.And([]comment.Matcher{
			comment.CreatedAfter{p.startDate},
			comment.BotMessage{p.keyword},
		})
	} else {
		return comment.BotMessage{p.keyword}
	}
}

func (p *Pinger) maxReached(pings []*github.IssueComment) bool {
	return p.maxCount != 0 && len(pings) >= p.maxCount
}

func (p *Pinger) shouldPingNow(pings []*github.IssueComment) bool {
	// We have never pinged, and we don't know when it started, ping now
	if len(pings) == 0 && p.startDate == nil {
		return true
	}

	var lastEvent *time.Time
	if len(pings) != 0 {
		lastEvent = pings[len(pings)-1].CreatedAt
	} else {
		lastEvent = p.startDate
	}

	return time.Since(*lastEvent) > p.timePeriod
}
