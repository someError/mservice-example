package main

func main() {
	pf := &priceFetcher{}
	service := NewLoggerService(NewMetricsService(pf))

	server := NewJSONApiServer(":3000", service)
	server.Run()
}
