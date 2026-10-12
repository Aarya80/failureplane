.PHONY: test race vet fmt check demo db-up db-down spike

test:
	go test ./...

race:
	go test -race ./...

vet:
	go vet ./...

fmt:
	gofmt -w ./cmd ./internal

check: test race vet spike

demo:
	go run ./cmd/failureplane demo

spike:
	python3 examples/faults/ambiguous_effect_spike.py

db-up:
	docker compose up -d postgres

db-down:
	docker compose down
