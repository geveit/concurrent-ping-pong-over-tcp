ping:
	@cd ping && go run .

pong:
	@cd pong && go run .

bench:
	@cd benchmarks && go test -bench=.

.PHONY: ping pong bench