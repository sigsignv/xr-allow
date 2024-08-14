xr-allow: *.go
	go build -o $@

.PHONY: clean
clean:
	$(RM) xr-allow

.PHONY: test
test:
	go test
