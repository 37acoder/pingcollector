package service

import "context"

type PingDomainSource interface {
	GetDomainList(ctx context.Context) []string
}

type ConstDomainSource struct {
}

func (c ConstDomainSource) GetDomainList(ctx context.Context) []string {
	return []string{
		"www.baidu.com",
		"www.google.com",
	}
}
