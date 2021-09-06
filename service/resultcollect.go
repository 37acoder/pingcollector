package service

import (
	"context"
	"fmt"

	"github.com/37acoder/pingcollector/ttype"
)

type PingResultManager interface {
	ReturnResponse(ctx context.Context, result ttype.PingResult)
}

type PrintPingResultReceiver struct {
}

func (i PrintPingResultReceiver) ReturnResponse(ctx context.Context, result ttype.PingResult) {
	fmt.Println(result)
}
