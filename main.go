// main.go
package main

import (
	"fmt"
	"log"

	orderpayment "github.com/akharb1192/orderpayment/ecommerce" // replace with your actual package path
)

// Mock implementations for PaymentGateway and InventoryService can be used during testing
// In this example, we use real implementations

// Real implementation of the PaymentGateway
type RealPaymentGateway struct{}

func (g *RealPaymentGateway) ProcessPayment(orderID string, amount float64) (bool, error) {
	// For the sake of simplicity, let's assume payment is always successful
	// In a real system, this would communicate with a payment processor like Stripe or PayPal.
	fmt.Printf("Processing payment for order %s, amount: %.2f\n", orderID, amount)
	return true, nil
}

// Real implementation of the InventoryService
type RealInventoryService struct{}

func (s *RealInventoryService) CheckStock(productID string) (int, error) {
	// For the sake of simplicity, assume we always have 10 items in stock
	// In a real system, this would query a database or an inventory management system.
	fmt.Printf("Checking stock for product %s\n", productID)
	return 10, nil
}

func main() {
	// Initialize real services
	paymentGateway := &RealPaymentGateway{}
	inventoryService := &RealInventoryService{}

	// Create an instance of the OrderService with the real dependencies
	orderService := orderpayment.NewOrderService(paymentGateway, inventoryService)

	// Place an order
	orderID := "order123"
	productID := "product456"
	quantity := 2
	amount := 200.00

	// Call the PlaceOrder method
	result, err := orderService.PlaceOrder(orderID, productID, quantity, amount)
	if err != nil {
		log.Fatalf("Error placing order: %v", err)
	}

	// Print the result of the order placement
	fmt.Println(result)
}
