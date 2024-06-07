package timestamp

import "time"

func GetTimestamp() int64 {
	return time.Now().UTC().UnixNano()
}
