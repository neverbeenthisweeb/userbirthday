package common

import (
	"context"

	uuid "github.com/satori/go.uuid"
)

type CtxKey string

const (
	CtxKeyRequestID CtxKey = "request_id"
)

func ContextWithRequestID() context.Context {
	return context.WithValue(context.Background(), CtxKeyRequestID, uuid.NewV4().String())
}
