package main

func main() {
	// client := client.NewClient("http://localhost:3000")

	// price, err := client.FetchPrice(context.Background(), "DASH")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(price.Price)
	// fmt.Println(price.Ticker)

	// return

	pf := &priceFetcher{}
	service := NewLoggerService(NewMetricsService(pf))

	server := NewJSONApiServer(":3000", service)
	server.Run()
}
