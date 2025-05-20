package repoUtils

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// TimestampToTimePtr converts a protobuf Timestamp to a *time.Time
func TimestampToTimePtr(ts *timestamppb.Timestamp) *time.Time {
	if ts == nil {
		return nil
	}

	t := ts.AsTime()

	return &t
}
