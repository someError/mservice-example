build:
	go build -o bin/price-fetcher

run: build
	./bin/price-fetcher