package middleware

import "category-service/pkg/llog"

type Middleware struct {
	fileLogger llog.Logger
}

func New(fileLogger llog.Logger) *Middleware {
	return &Middleware{
		fileLogger: fileLogger,
	}
}
