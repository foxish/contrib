package states

import "k8s.io/contrib/mungegithub/github"

type PreReview struct{}

func Process(obj *github.MungeObject) (*State, error) {
	return nil
}

func Name() string {
	return "PreReview"
}
