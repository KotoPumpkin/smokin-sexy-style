package main

import "fmt"

type Item struct {
	Name  string
	Price float64
	Stock int
}

type StockManager interface {
	CheckStock() int
	UpdateStock(quantity int)
	PrintStockInfo()
}

type Electronics struct {
	Item
	Brand string
	Model string
}

func (e *Electronics) CheckStock() int {
	return e.Stock
}

func (e *Electronics) UpdateStock(quantity int) {
	e.Stock += quantity
}

func (e *Electronics) PrintStockInfo() {
	fmt.Printf("Brand: %s, Model: %s\n", e.Brand, e.Model)
	fmt.Printf("Stock: %d\n", e.Stock)
}

func main() {
	iphone := &Electronics{
		Item: Item{
			Name:  "羽中的亲妈",
			Price: 999.99,
			Stock: 10,
		},
		Brand: "雷蒙",
		Model: "雷蒙太太",
	}

	fmt.Println("Initial Stock:")
	iphone.PrintStockInfo()

	fmt.Println("Checking stock...")
	quantity := iphone.CheckStock()
	fmt.Printf("Current stock: %d\n", quantity)

	fmt.Println("Updating stock...")
	iphone.UpdateStock(5)

	fmt.Println("Updated Stock:")
	iphone.PrintStockInfo()
}
