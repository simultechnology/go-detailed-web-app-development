package main

import (
	"context"
	"fmt"
)

type TraceID string

const ZeroTraceID = ""

type traceIDKey struct{}

func SetTraceID(ctx context.Context, traceId TraceID) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceId)
}

func GetTraceID(ctx context.Context) TraceID {
	if v, ok := ctx.Value(traceIDKey{}).(TraceID); ok {
		return v
	}
	return ZeroTraceID
}

func main() {
	fmt.Println("ch2-7 starts!")

	ctx := context.Background()
	fmt.Printf("trace id = %q\n", GetTraceID(ctx))
	ctx = SetTraceID(ctx, "test-id")
	fmt.Printf("trace id = %q\n", GetTraceID(ctx))
}
