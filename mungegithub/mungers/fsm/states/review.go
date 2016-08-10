package states

import "k8s.io/contrib/mungegithub/github"

type Review struct{}

func Process(obj *github.MungeObject) (*State, error) {
	return nil
}

func Name() string {
	return "Review"
}
