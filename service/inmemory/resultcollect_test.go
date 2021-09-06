package inmemory

import (
	"fmt"
	"testing"
	"time"
)

func TestSomething(t *testing.T) {
	r := NewResultCollector()
	go func() {
		for range time.Tick(time.Second) {
			r.AddDataPoint("domain", DataPoint{
				time: time.Now(),
				RTL:  1000,
			})
		}
	}()
	count := 0
	for range time.Tick(time.Millisecond * 100) {
		fmt.Println(r.CalcDomainAverage("domain"))
		count++
		if count == 100 {
			return
		}
	}
}
