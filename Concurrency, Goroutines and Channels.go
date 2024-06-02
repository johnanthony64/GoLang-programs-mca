package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Restaurant struct {
	menu map[string]int // Menu items with their available quantities
	mu   sync.Mutex     // Mutex to synchronize access to the menu
}

func NewRestaurant() *Restaurant {
	return &Restaurant{
		menu: map[string]int{
			"Pizza":     10,
			"Burger":    15,
			"Pasta":     20,
			"Salad":     5,
			"Sandwich":  10,
			"Soft Drink": 50,
		},
	}
}

func (r *Restaurant) OrderItem(item string, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()

	r.mu.Lock()
	defer r.mu.Unlock()

	fmt.Printf("Customer is trying to order %s\n", item)

	if quantity, ok := r.menu[item]; ok && quantity > 0 {
		// Simulating order processing delay
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		r.menu[item]--
		ch <- fmt.Sprintf("Ordered %s", item)
		fmt.Printf("Customer ordered %s\n", item)
	} else {
		ch <- fmt.Sprintf("%s is not available", item)
		fmt.Printf("%s is not available\n", item)
	}
}

func main() {
	restaurant := NewRestaurant()

	// Create a channel to communicate order statuses
	orderCh := make(chan string)

	// WaitGroup to synchronize goroutines
	var wg sync.WaitGroup

	// Simulate multiple customers ordering items concurrently
	for _, item := range []string{"Pizza", "Burger", "Pasta", "Salad", "Pizza", "Soft Drink", "Sandwich"} {
		wg.Add(1)
		go restaurant.OrderItem(item, &wg, orderCh)
	}

	// Close the channel after all orders are done
	go func() {
		wg.Wait()
		close(orderCh)
	}()

	// Print order statuses
	for order := range orderCh {
		fmt.Println(order)
	}
}
