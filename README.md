# Kinso Interview

Kinso Backend Engineering Challenge - Option B: Data Pipeline Focus.

High-level overview can be found in here: https://docs.google.com/document/d/174QbVohxZXIcQhpWFaaGUa-VAyY1u_RS19MGOpndegQ/edit?usp=sharing

# Requirements 

- [Go](https://go.dev/dl/)

# Documentation

The app has main runners:
- `cmd/generator/main.go` - to mock event data (creates 100 events for 20 users) in the `_data/events.json` file.
- `cmd/app/main.go` - actual service to handle generated events

App when started (see useful commands) runs on :8080 port.

All the logic is inside of the `app/`.

As described in the high-level overview, the idea is to create a queue listener (`app/listener/listener.go`) in our simplified case I made a trade off to poll the `_data/events.json` every 10 seconds.

If listener found events, then it calls the consumer (`app/consumer/consumer.go`) which is responsible for figuring out the priority of an event based on the configuration in the `app/consumer/config.go` and then storing it in DB (`app/storage/storage.go`).

DB has two indexes:
- primary hash by event ID
- and user ID hash with entries order by weight (priority) and timestamp in case of the tie.

App exposes an endpoint to read ordered events for a given user ID http://localhost:8080/api/v1/events/<user_id>, user ID can be picked up from `_data/events.json` file.

# Useful comands
- Run the app `make run` or `go run cmd/app/main.go`
- Run the event data generator `make gen` `go run cmd/generator/main.go`