APP      := xr-allow
LD_FLAGS := "-w -s"

.PHONY: \
	xr-allow \
	linux-amd64 \
	linux-arm64 \
	darwin-amd64 \
	darwin-arm64 \
	windows-amd64 \
	build \
	clean \
	dev \
	test

xr-allow: *.go
	@ go build -v -o $(APP) ./...

linux-amd64:
	@ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags=$(LD_FLAGS) -o $(APP).$@

linux-arm64:
	@ CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags=$(LD_FLAGS) -o $(APP).$@

darwin-amd64:
	@ GOOS=darwin GOARCH=amd64 go build -ldflags=$(LD_FLAGS) -o $(APP).$@

darwin-arm64:
	@ GOOS=darwin GOARCH=arm64 go build -ldflags=$(LD_FLAGS) -o $(APP).$@

windows-amd64:
	@ GOOS=windows GOARCH=amd64 go build -ldflags=$(LD_FLAGS) -o $(APP).$@.exe

build: linux-amd64 linux-arm64 darwin-amd64 darwin-arm64 windows-amd64

clean:
	@ $(RM) $(APP) $(APP).*

test:
	go test -v ./...
