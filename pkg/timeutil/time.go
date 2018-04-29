package timeutil

import (
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
)

// TimeToTimestamp converts a Go time object into a proto timestamp
func TimeToTimestamp(t time.Time) *timestamp.Timestamp {
	if t.IsZero() {
		return nil
	}

	return &timestamp.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Sub(time.Unix(t.Unix(), 0))),
	}
}

// TimestampToTime converts a proto timestamp to a Go time object
func TimestampToTime(ts *timestamp.Timestamp) time.Time {
	return time.Unix(ts.Seconds, int64(ts.Nanos)).UTC()
}
