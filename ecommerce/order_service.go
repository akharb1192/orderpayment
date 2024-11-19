// order_service.go
package ecommerce

import "fmt"

// PaymentGateway defines the method to process payments.
type PaymentGateway interface {
	ProcessPayment(orderID string, amount float64) (bool, error)
}

// InventoryService defines the method to check product stock.
type InventoryService interface {
	CheckStock(productID string) (int, error)
}

// OrderService handles the logic for placing and managing orders.
type OrderService struct {
	paymentGateway   PaymentGateway
	inventoryService InventoryService
}

// NewOrderService creates a new instance of OrderService.
func NewOrderService(paymentGateway PaymentGateway, inventoryService InventoryService) *OrderService {
	return &OrderService{
		paymentGateway:   paymentGateway,
		inventoryService: inventoryService,
	}
}

// PlaceOrder places an order and processes the payment.
func (s *OrderService) PlaceOrder(orderID, productID string, quantity int, amount float64) (string, error) {
	// Check stock
	stock, err := s.inventoryService.CheckStock(productID)
	if err != nil {
		return "", err
	}
	if stock < quantity {
		return "", fmt.Errorf("insufficient stock")
	}

	// Process payment
	paymentSuccess, err := s.paymentGateway.ProcessPayment(orderID, amount)
	if err != nil {
		return "", err
	}
	if !paymentSuccess {
		return "", fmt.Errorf("payment failed")
	}

	// Place the order (here we just return a success message)
	return "Order placed successfully", nil
}
