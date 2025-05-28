package utils

import (
	"log/slog"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func TimeNow() *time.Time {
	now := time.Now()
	return &now
}

func MustToTime(tStr string) *time.Time {
	t, err := time.Parse(time.RFC3339, tStr)
	if err != nil {
		slog.Warn(err.Error())
		return nil
	}

	return &t
}

func MustToTimestamp(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}

	return timestamppb.New(*t)
}
