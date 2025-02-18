package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/eliseomartelli/micromachine"
)

type OrderState string

const (
	OrderCreated   OrderState = "created"
	OrderPaid      OrderState = "paid"
	OrderShipped   OrderState = "shipped"
	OrderDelivered OrderState = "delivered"
	OrderCancelled OrderState = "cancelled"
)

type Order struct {
	mu             sync.Mutex
	ID             int        `json:"id"`
	Status         OrderState `json:"status"`
	Items          []string   `json:"items"`
	TrackingNumber string     `json:"tracking_number"`
}

func orderProcessingExample() {
	order := &Order{
		ID:     123,
		Items:  []string{"Product A", "Product B"},
		Status: OrderCreated,
	}

	machine := micromachine.NewMicromachine(order.Status)

	machine.AddTransition(OrderCreated, OrderPaid, func() error {
		order.mu.Lock()
		defer order.mu.Unlock()
		fmt.Println("Order paid. Processing payment...")
		order.Status = OrderPaid
		return nil
	})

	machine.AddTransition(OrderPaid, OrderShipped, func() error {
		order.mu.Lock()
		defer order.mu.Unlock()
		fmt.Println("Order shipped. Generating tracking number...")
		order.Status = OrderShipped
		order.TrackingNumber = "ABC123XYZ"
		return nil
	})

	machine.AddTransition(OrderShipped, OrderDelivered, func() error {
		order.mu.Lock()
		defer order.mu.Unlock()
		fmt.Println("Order delivered. Updating inventory...")
		order.Status = OrderDelivered
		return nil
	})

	machine.AddTransition(OrderCreated, OrderCancelled, func() error {
		order.mu.Lock()
		defer order.mu.Unlock()
		fmt.Println("Order cancelled.")
		order.Status = OrderCancelled
		return nil
	})

	machine.AddTransition(OrderPaid, OrderCancelled, func() error {
		order.mu.Lock()
		defer order.mu.Unlock()
		fmt.Println("Order cancelled.")
		order.Status = OrderCancelled
		return nil
	})

	http.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		order.mu.Lock()
		defer order.mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(order)
	})

	http.HandleFunc("/transition", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			To OrderState `json:"to"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := machine.Transition(req.To); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	fmt.Println("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func main() {
	orderProcessingExample()
}
