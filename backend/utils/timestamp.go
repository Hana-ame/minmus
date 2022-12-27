package utils

import (
	"sync"
	"time"
)

var DefaultTS = NewTS()

type TS struct {
	last int64
	inc  int64
	sync.Mutex
}

func NewTS() *TS {
	return &TS{
		last: getTimeStamp(),
		inc:  1,
	}
}

func (ts *TS) GetTS() int64 {
	ts.Lock()
	defer ts.Unlock()
	now := getTimeStamp()
	if now != ts.last {
		// if ts.last+ts.inc < now
		ts.last = now
		ts.inc = 0
	} else {
		ts.inc++
		now += ts.inc
	}
	return now
}

func getTimeStamp() int64 {
	return time.Now().UnixMilli() << 16
}

func GetTS() int64 {
	return DefaultTS.GetTS()
}
