package service

import (
	"context"
	"log"

	"github.com/37acoder/pingcollector/ttype"
	"github.com/37acoder/pingcollector/utils"
)

type PingService interface {
	StartPingService(ctx context.Context, pingDomainSource PingDomainSource, receiver PingResultManager) error
}

type DefaultPingService struct {
	ResultCollector chan ttype.PingResult
}

func NewDefaultPingService() *DefaultPingService {
	return &DefaultPingService{ResultCollector: make(chan ttype.PingResult, 0)}
}

func (d *DefaultPingService) StartPingService(ctx context.Context, pingDomainSource PingDomainSource, receiver PingResultManager) error {
	PingerMap := map[string]*utils.Pinger{}
	go func() {
		for {
			select {
			case <-ctx.Done():
				for domain, pinger := range PingerMap {
					pinger.Stop()
					log.Println(domain, "stop ping loop")
				}
				return
			case result := <-d.ResultCollector:
				receiver.ReturnResponse(ctx, result)
			}
		}
	}()

	for _, domain := range pingDomainSource.GetDomainList(ctx) {
		pinger, err := utils.NewPinger(domain, d.ResultCollector)
		if err != nil || pinger == nil {
			return err
		}
		PingerMap[domain] = pinger
	}
	for domain, pinger := range PingerMap {
		pinger.Start()
		log.Println(domain, "start ping loop")
	}
	return nil
}
