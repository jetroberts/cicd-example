package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type TimeEvent struct {
	Dates []time.Time `json:"dates"`
}

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, event TimeEvent) (string, error) {
	if len(event.Dates) < 2 {
		err := errors.New("at least two dates must be provided")
		return "", err
	}

	return fmt.Sprintf("There have been %v hours between %v and %v", event.Dates[0].Sub(event.Dates[1]).Hours(), event.Dates[0].Format(time.RFC1123), event.Dates[1].Format(time.RFC1123)), nil
}
