package serializer

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func PbTimestampToTime(inputProto *timestamppb.Timestamp) *time.Time {
	if inputProto != nil {
		outputTime := inputProto.AsTime()

		return &outputTime
	}

	return nil
}

// TimeToPbTimestamp sends out times without timezone.
func TimeToPbTimestamp(ip *time.Time) *timestamppb.Timestamp {
	if ip != nil {
		timeWithoutTZ := time.Date(ip.Year(), ip.Month(), ip.Day(), ip.Hour(), ip.Minute(), ip.Second(), ip.Nanosecond(), time.UTC)

		return timestamppb.New(timeWithoutTZ)
	}

	return nil
}
