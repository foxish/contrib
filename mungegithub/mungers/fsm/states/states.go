package states

import "k8s.io/contrib/mungegithub/github"

func ProcessCurrentState(obj *github.MungeObject) error {
	currentState := PreReview{obj}

	for {
		currentState, err = currentState.Process(obj)
		if err != nil {
			return err
		}
		if currentState == nil {
			break
		}
	}

	return nil
}

type State interface {
	Process(obj *github.MungeObject) (*State, error)
	Name() string
}
