package main

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_HandleRequest(t *testing.T) {
	tests := []struct {
		name           string
		testDates      TimeEvent
		expectedResult string
		expectedError  bool
	}{
		{
			name: "First date greater than second date",
			testDates: TimeEvent{
				Dates: []time.Time{
					time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
					time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			expectedResult: "There have been 8040 hours between Mon, 02 Jan 2023 00:00:00 UTC and Tue, 01 Feb 2022 00:00:00 UTC",
		},
		{
			name: "Second date greater than first date ",
			testDates: TimeEvent{
				Dates: []time.Time{
					time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			},
			expectedResult: "There have been -8040 hours between Tue, 01 Feb 2022 00:00:00 UTC and Mon, 02 Jan 2023 00:00:00 UTC",
		},
		{
			name: "Three dates, still looks at first two",
			testDates: TimeEvent{
				Dates: []time.Time{
					time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
					time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			},
			expectedResult: "There have been -8040 hours between Tue, 01 Feb 2022 00:00:00 UTC and Mon, 02 Jan 2023 00:00:00 UTC",
		},
		{
			name: "One dates throws an error",
			testDates: TimeEvent{
				Dates: []time.Time{
					time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			result, err := HandleRequest(ctx, tt.testDates)
			if err != nil && tt.expectedError {
				return
			}

			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, tt.expectedResult, result)
		})
	}
}