# Datadog Agent (Docker) + Go Sample App

Example of running a [Datadog](https://www.datadoghq.com/ "Datadog") Agent in a docker cotainer and send custom metrics from a Go sample application using the Go statsd library.

## How to run it

1. You need an API Key from your account in Datadog. You can get one [here](https://app.datadoghq.com/account/settings#api "here").

2. Create `.env`file in the root folder with the following content:

	`DD_API_KEY=xxxx`

3. Run datadog agent as a docker container:

	`docker-compose up -d`

4. Install the Dogstatsd client:

	`go get github.com/DataDog/datadog-go/statsd`

5. Execute go application:

	`go run main.go`
