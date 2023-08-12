package common

import (
	"context"
	"fmt"
	"time"

	"incrowd-backend/internal/context_wrapper"
	"incrowd-backend/internal/log"
)

// ConvertStringToDate converts a string representation of a date into a time.Time value using the specified layout
func ConvertStringToDate(ctx context.Context, layout, dateStr string) time.Time {
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		log.Errort(context_wrapper.GetCorrelationID(ctx), fmt.Sprintf("could not convert string '%s' to a date with layout '%s'", dateStr, layout))
		return time.Now()
	}

	return date
}
