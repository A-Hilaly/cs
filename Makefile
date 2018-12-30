build:
	GO111MODULE=on \
		go build -o cs ./cmd/main.go

test:
	GO111MODULE=on \
		go test ./pkg/...

test-cov:
	GO111MODULE=on \
		go test -cover ./pkg/...

benchmark:
	

.PHONY: all test
