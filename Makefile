.PHONY: test ut it
all: ut

ut:
	@echo "unit testing..."
	go test -v -timeout 30s .

it: ut
	@echo "integration testing..."
	go test -v . -tags=integration -args ${params}

test: it
