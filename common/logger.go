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

func LogWarn(ctx context.Context, msg string, err error) {
	var requestID string
	id, ok := ctx.Value(CtxKeyRequestID).(string)
	if ok {
		requestID = id
	}
	warnMsg := fmt.Sprintf("[%s] %s WARN: %s",
		time.Now().Format(time.RFC822),
		requestID,
		msg,
	)
	if err != nil {
		warnMsg += fmt.Sprintf(" -> err=%s", err.Error())
	}
	warnMsg += "\n"
	fmt.Print(warnMsg)
}

func LogErr(ctx context.Context, msg string, err error) {
	var requestID string
	id, ok := ctx.Value(CtxKeyRequestID).(string)
	if ok {
		requestID = id
	}
	fmt.Printf("[%s] %s ERROR: %s -> err=%s\n",
		time.Now().Format(time.RFC822),
		requestID,
		msg,
		err.Error(),
	)
}
