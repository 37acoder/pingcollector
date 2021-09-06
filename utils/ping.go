package utils

import (
	"log"

	"github.com/37acoder/pingcollector/ttype"
	"github.com/go-ping/ping"
)

type Pinger struct {
	pinger *ping.Pinger
	Domain string
}

func NewPinger(domain string, resultReceiver chan<- ttype.PingResult) (*Pinger, error) {
	pinger, err := ping.NewPinger(domain)
	if err != nil {
		return nil, err
	}
	pinger.OnRecv = func(packet *ping.Packet) {
		resultReceiver <- ttype.PingResult{
			Domain:    domain,
			PingDelay: packet.Rtt,
		}
	}
	return &Pinger{
		pinger: pinger,
		Domain: domain,
	}, nil
}

func (p *Pinger) Stop() {
	p.pinger.Stop()
}

func (p *Pinger) Start() {
	go func() {
		err := p.pinger.Run()
		if err != nil {
			log.Println("start pinger error", err)
			return
		}
	}()
}
