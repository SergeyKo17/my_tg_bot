package handlers

import (
	"context"
)

type Handlers struct{}

func (h Handlers) Info(ctx context.Context) string {
	return "My name is Sergey"
}

func (h Handlers) GitLink(ctx context.Context) string {
	return "My Git's link"
}

func (h Handlers) MainHandler(mainContext context.Context, text string) string {
	var res string
	ctx, cancel := context.WithCancel(mainContext)
	defer cancel()

	switch text {
	case "Info":
		res = h.Info(ctx)
	case "GitLink":
		res = h.GitLink(ctx)
	}

	return res
}
