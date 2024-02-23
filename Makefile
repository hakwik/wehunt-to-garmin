BINARY_NAME=wehunt-to-garmin
LDFLAGS="-s -w"

all:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -ldflags ${LDFLAGS} -o out/${BINARY_NAME}-macos-intel64 .
	CGO_ENABLED=0 GOARCH=arm64 GOOS=darwin go build -ldflags ${LDFLAGS} -o out/${BINARY_NAME}-macos-arm64 .
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags ${LDFLAGS} -o out/${BINARY_NAME}-linux64 .
	CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build -ldflags ${LDFLAGS} -o out/${BINARY_NAME}-windows64 .

clean:
	go clean
	rm -rf out