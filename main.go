package main

import (
	"fmt"
	"github.com/datadog-go/statsd"
	"time"
)

func main() {
	client, err := statsd.New("localhost:8125",
		statsd.WithNamespace("custom_metric."), // prefix every metric with the app name
	)

	if err != nil {
		fmt.Printf("%v", err.Error())
		panic(1)
	}

	for {
		err := client.Incr("test", nil, 1)
		fmt.Print(".")

		if err != nil {
			fmt.Printf("%v", err.Error())
		}

		time.Sleep(1 * time.Second)
	}
}
