package common

import (
	"context"
	"fmt"
	"time"
)

func LogInfo(ctx context.Context, msg string) {
	var requestID string
	id, ok := ctx.Value(CtxKeyRequestID).(string)
	if ok {
		requestID = id
	}
	fmt.Printf("[%s] %s INFO: %s\n",
		time.Now().Format(time.RFC822),
		requestID,
		msg,
	)
}

func LogErr(ctx context.Context, err error) {
	var requestID string
	id, ok := ctx.Value(CtxKeyRequestID).(string)
	if ok {
		requestID = id
	}
	fmt.Printf("[%s] %s ERROR: %s\n",
		time.Now().Format(time.RFC822),
		requestID,
		err.Error(),
	)
}
