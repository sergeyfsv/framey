
.PHONY: install test test-coverage

default: install


install:  ## fetch and install all required deps
	go get -u github.com/Masterminds/log-go@v0.4.0
	go get -u go.jonnrb.io/speedtest"v0.2.0
	go get -u golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
)

test:
	go test -v ./... -short


# Generate test coverage
test-cover:
	go test -v -coverprofile cover.out .