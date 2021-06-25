GO=go
fib:
	$(GO) build -o fib server.go
	./fib

test:
	$(GO) test -bench .
	$(GO) test -bench .

.PHONY: clean

clean:
	rm -f fib
