package inmemory

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/37acoder/pingcollector/ttype"
)

//StoreMaximum 24小时*7
//const StoreMaximum = 1 * 60 * 60 * 24 * 7
const StoreMaximum = 5

type ResultCollector struct {
	metrics map[string]*Series
	sync.RWMutex
}

func NewResultCollector() *ResultCollector {
	return &ResultCollector{metrics: map[string]*Series{}}
}

func (r *ResultCollector) GetSeriesByDomain(domain string) *Series {
	r.RLock()
	s := r.metrics[domain]
	r.RUnlock()
	return s
}

func (r *ResultCollector) AddDataPoint(domain string, point DataPoint) {
	var series *Series
	r.RLock()
	if s, ok := r.metrics[domain]; ok {
		series = s
	}
	r.RUnlock()
	if series == nil {
		r.Lock()
		if s, ok := r.metrics[domain]; ok {
			series = s
		} else {
			series = &Series{}
			r.metrics[domain] = series
		}
		r.Unlock()
	}
	series.AddDataPoint(point)
}

type Series struct {
	data   [StoreMaximum]DataPoint
	start  int64
	length int64
	sync.RWMutex
}

func (s *Series) LastIndex() int64 {
	return (s.start+s.length)%StoreMaximum - 1
}

func (s *Series) AddDataPoint(point DataPoint) {
	s.Lock()
	newDataPointIndex := (s.start + s.length) % StoreMaximum
	s.data[newDataPointIndex] = point
	if newDataPointIndex == s.start {
		s.start += 1
	}
	if s.length != StoreMaximum {
		s.length += 1
	}
	s.Unlock()
}

func (s *Series) Iter(f func(point DataPoint, series *Series)) {
	s.RLock()
	for count := int64(0); count != s.length; count++ {
		dataPoint := s.data[count%StoreMaximum]
		f(dataPoint, s)
	}
	s.RUnlock()
}

type DataPoint struct {
	time time.Time
	RTL  time.Duration
}

func (r *ResultCollector) ReturnResponse(ctx context.Context, result ttype.PingResult) {
	r.AddDataPoint(result.Domain, DataPoint{
		time: time.Now(),
		RTL:  result.PingDelay,
	})
}

func (r *ResultCollector) CalcDomainAverage(domain string) (time.Duration, error) {
	s := r.GetSeriesByDomain(domain)
	if s == nil {
		return 0, fmt.Errorf("series %s not found", domain)
	}
	var avg float64
	s.Iter(func(point DataPoint, series *Series) {
		fmt.Println(point.RTL, series.length)
		avg += float64(point.RTL) / float64(series.length)
	})
	return time.Duration(avg), nil
}
