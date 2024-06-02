package main

import (
	"fmt"
)

// Restaurant information
var restaurantName = "Shiro"
var restaurantLocation = "14, Vittal Malya Road, Bengaluru -01"

// Menu items with prices 
var menuItems = map[string]float64{
	"Chicken": 199.00,
	"Paneer":  229.00,
	"Dal":     149.00,
	"Biryani": 179.00,
	"Naan":    59.00,
	"Raita":   49.00,
}

// Tax rate
const taxRate = 0.10

// Discount code and percentage
var discountCode = "SPICY15"
var discountPercentage = 0.15

// Customer order represented as a struct
type Order struct {
	name      string
	items     map[string]int // Map to store item and its quantity
	totalCost float64
}

func main() {
	// Display restaurant information
	fmt.Println("Welcome to", restaurantName, "at", restaurantLocation)

	// Display menu
	fmt.Println("\nMenu:")
	for itemName, price := range menuItems {
		fmt.Printf("- %-20s: ₹%.2f\n", itemName, price)
	}

	// Create a new customer order
	var order Order
	order.items = make(map[string]int)
	fmt.Println("\nPlease enter your name:")
	fmt.Scanln(&order.name)

	// Order loop
	for {
		var item string
		fmt.Println("\nWhat would you like to order? (Enter 'q' to quit):")
		fmt.Scanln(&item)

		if item == "q" {
			break
		}

		// Validate item selection (check if item exists in menu)
		if _, ok := menuItems[item]; !ok {
			fmt.Println("Invalid item. Please choose from the menu:")
			continue
		}

		// Add quantity
		var quantity int
		fmt.Println("Enter quantity:")
		fmt.Scanln(&quantity)

		// Add item to order and update total cost
		order.items[item] += quantity
		order.totalCost += menuItems[item] * float64(quantity)

		fmt.Printf("Added %d %s to your order (Total cost: ₹%.2f)\n", quantity, item, order.totalCost)
	}

	// Apply tax
	order.totalCost *= (1 + taxRate)

	// Apply discount if valid code entered
	var enteredCode string
	fmt.Println("\nDo you have a discount code? (Enter 'y' or 'n'):")
	fmt.Scanln(&enteredCode)

	if enteredCode == "y" {
		fmt.Println("Enter your discount code:")
		fmt.Scanln(&enteredCode)

		if enteredCode == discountCode {
			fmt.Println("Discount applied!")
			order.totalCost *= (1 - discountPercentage)
		} else {
			fmt.Println("Invalid discount code.")
		}
	}

	// Display order summary
	if len(order.items) > 0 {
		fmt.Println("\nYour order summary:")
		fmt.Println("Customer:", order.name)
		fmt.Println("Items:")
		for item, quantity := range order.items {
			fmt.Printf("- %d %s\n", quantity, item)
		}
		fmt.Printf("Total cost (after tax and discount): ₹%.2f\n", order.totalCost)
		
		// Order confirmation
		var confirm string
		fmt.Println("Confirm your order? (Enter 'yes' to confirm, any other key to cancel)")
		fmt.Scanln(&confirm)
		if confirm == "yes" {
			fmt.Println("Order confirmed! Thank you for your order!")
		} else {
			fmt.Println("Order canceled.")
		}
	} else {
		fmt.Println("You haven't ordered anything. Please come back again soon!")
	}
}
