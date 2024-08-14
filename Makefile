xr-allow: *.go
	go build -o $@

.PHONY: clean
clean:
	$(RM) xr-allow
