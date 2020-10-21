.PHONY: test ut it
all: ut

ut:
	@echo "unit testing..."
	go test -timeout 30s .

it: ut
	@echo "integration testing..."
	go test -timeout 30s . -tags=integration -args ${params}

test: it
