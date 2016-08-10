package ui

import (
	"fmt"

	mgh "k8s.io/contrib/mungegithub/github"
)

func Comment(
	obj *mgh.MungeObject,
	action string,
	args string,
	description string,
) error {
	return obj.WriteComment(
		fmt.Sprintf(
			"[%s] %s\n\n%s",
			action,
			args,
			description,
		),
	)

}
