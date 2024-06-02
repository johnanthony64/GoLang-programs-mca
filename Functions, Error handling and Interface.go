package main

import (
    "errors"
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

// Order interface
type Order interface {
    AddItem(string, int) error
    CalculateTotalCost() float64
    ApplyTax()
    ApplyDiscount(string) error
    GetOrderSummary() string
}

// Order struct
type OrderImpl struct {
    name      string
    items     map[string]int
    totalCost float64
}

// AddItem method
func (o *OrderImpl) AddItem(item string, quantity int) error {
    if _, ok := menuItems[item]; !ok {
        return errors.New("invalid item")
    }
    o.items[item] += quantity
    o.totalCost += menuItems[item] * float64(quantity)
    return nil
}

// CalculateTotalCost method
func (o *OrderImpl) CalculateTotalCost() float64 {
    return o.totalCost
}

// ApplyTax method
func (o *OrderImpl) ApplyTax() {
    o.totalCost *= (1 + taxRate)
}

// ApplyDiscount method
func (o *OrderImpl) ApplyDiscount(code string) error {
    if code == discountCode {
        o.totalCost *= (1 - discountPercentage)
        return nil
    }
    return errors.New("invalid discount code")
}

// GetOrderSummary method
func (o *OrderImpl) GetOrderSummary() string {
    summary := fmt.Sprintf("Customer: %s\nItems:\n", o.name)
    for item, quantity := range o.items {
        summary += fmt.Sprintf("- %d %s\n", quantity, item)
    }
    summary += fmt.Sprintf("Total cost (after tax and discount): ₹%.2f\n", o.totalCost)
    return summary
}

// DisplayMenu function
func DisplayMenu() {
    fmt.Println("Welcome to", restaurantName, "at", restaurantLocation)
    fmt.Println("\nMenu:")
    for itemName, price := range menuItems {
        fmt.Printf("- %-20s: ₹%.2f\n", itemName, price)
    }
}

func main() {
    DisplayMenu()

    // Create a new customer order
    var order OrderImpl
    order.items = make(map[string]int)

    fmt.Println("\nPlease enter your name:")
    fmt.Scanln(&order.name)

    for {
        var item string
        fmt.Println("\nWhat would you like to order? (Enter 'q' to quit):")
        fmt.Scanln(&item)
        if item == "q" {
            break
        }

        // Add quantity
        var quantity int
        fmt.Println("Enter quantity:")
        fmt.Scanln(&quantity)

        // Add item to order
        err := order.AddItem(item, quantity)
        if err != nil {
            fmt.Println(err)
            continue
        }

        fmt.Printf("Added %d %s to your order (Total cost: ₹%.2f)\n", quantity, item, order.CalculateTotalCost())
    }

    // Apply tax
    order.ApplyTax()

    // Apply discount
    var enteredCode string
    fmt.Println("\nDo you have a discount code? (Enter 'y' or 'n'):")
    fmt.Scanln(&enteredCode)
    if enteredCode == "y" {
        fmt.Println("Enter your discount code:")
        fmt.Scanln(&enteredCode)
        err := order.ApplyDiscount(enteredCode)
        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Println("Discount applied!")
        }
    }

    // Display order summary
    if len(order.items) > 0 {
        fmt.Println("\nYour order summary:")
        fmt.Println(order.GetOrderSummary())

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