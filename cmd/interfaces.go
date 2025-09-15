package main

import "context"

type IHandlers interface {
	MainHandler(context.Context, string) string
}
