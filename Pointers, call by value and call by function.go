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
	ConfirmOrder() string
}

// Order struct
type OrderImpl struct {
	name      string
	items     map[string]int
	totalCost float64
	confirmed bool
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

// ConfirmOrder method
func (o *OrderImpl) ConfirmOrder() string {
	o.confirmed = true
	return "Order confirmed. Thank you for your purchase!"
}

// Function to modify the order name using pointers
func ModifyOrderName(order *OrderImpl, newName string) {
	order.name = newName
}

// Function to add an item to the order using call by value
func AddItemToOrder(order OrderImpl, item string, quantity int) (OrderImpl, error) {
	newOrder := order
	err := newOrder.AddItem(item, quantity)
	return newOrder, err
}

// Function to apply tax to the order using pointers
func ApplyTaxToOrder(order *OrderImpl) {
	order.ApplyTax()
}

// Function to apply discount to the order using call by value
func ApplyDiscountToOrder(order OrderImpl, code string) (OrderImpl, error) {
	err := order.ApplyDiscount(code)
	return order, err
}

// Function to get user input for confirmation
func ConfirmOrderInput() string {
	var confirmation string
	fmt.Println("Would you like to confirm your order? (yes/no)")
	fmt.Scanln(&confirmation)
	return confirmation
}

// Function to get user input for discount code
func GetDiscountCodeInput() string {
	var code string
	fmt.Println("Please enter your discount code:")
	fmt.Scanln(&code)
	return code
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
		fmt.Println("\nWhat would you like to order? (enter 'done' to finish)")
		fmt.Scanln(&item)

		if item == "done" {
			break
		}

		var quantity int
		fmt.Println("How many", item, "do you want?")
		fmt.Scanln(&quantity)

		newOrder, err := AddItemToOrder(order, item, quantity)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			order = newOrder
			fmt.Println("Item added to the order.")
		}
	}

	ApplyTaxToOrder(&order)

	discountCode := GetDiscountCodeInput()
	order, err := ApplyDiscountToOrder(order, discountCode)
	if err != nil {
		fmt.Println("Error applying discount:", err)
	}

	fmt.Println("\nOrder Summary:")
	fmt.Println(order.GetOrderSummary())

	confirmation := ConfirmOrderInput()
	if confirmation == "yes" {
		order.ConfirmOrder()
		fmt.Println(order.ConfirmOrder())
	} else {
		fmt.Println("Order not confirmed. Thank you for visiting!")
	}
}