package util

import "time"

func NowMillisecond() int64 {
	return time.Now().UnixNano()/int64(time.Millisecond)
}