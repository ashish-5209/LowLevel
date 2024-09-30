// +-------------------+       1           +-------------------+
// | VendingMachine    |<----------------->| Inventory         |
// |-------------------|                   |-------------------|
// | balance           |                   | items[]: Item     |
// | totalSales        |                   |-------------------|
// | currentSelection  |                   | +addItem(item)    |
// |-------------------|                   | +removeItem(item) |
// | +selectItem()     |                   | +getItemDetails() |
// | +insertMoney()    |                   | +updateStock()    |
// | +dispenseItem()   |                   +-------------------+
// | +refund()         |
// | +cancel()         |
// +-------------------+                         |
//       |                                        |
//       |                                        v
//       v                             +-------------------+
// +-------------------+               | Item              |
// | Payment           | 1 ----------->|-------------------|
// |-------------------|               | id: int           |
// | paymentType       |               | name: string      |
// | amountPaid        |               | price: float      |
// |-------------------|               | quantity: int     |
// | +acceptPayment()  |               |-------------------|
// | +processPayment() |               | +updateQuantity() |
// +-------------------+               +-------------------+
//       |                                        |
//       |                                        v
//       | 1                         +-------------------+
//       +-------------------------->| Coin              |
//       |                            |-------------------|
//       |                            | value: float      |
//       |                            |-------------------|
//       |                            | +getValue()       |
//       |                            +-------------------+
//       |
//       | 1                         +-------------------+
//       +-------------------------->| Card              |
//                                    |-------------------|
//                                    | cardNumber: string|
//                                    | expiryDate: string|
//                                    |-------------------|
//                                    | +chargeCard()     |
//                                    +-------------------+

// +-------------------+               +-------------------+
// | Transaction       |               | Display           |
// |-------------------|               |-------------------|
// | transactionId     |               | message: string   |
// | dateTime          |               |-------------------|
// | item: Item        |               | +showMessage()    |
// | amountPaid        |               | +displayPrice()   |
// | paymentStatus     |               +-------------------+
// |-------------------|
// | +createReceipt()  |
// | +recordTransaction|
// +-------------------+

package main

import (
	"fmt"
	"time"
)

// Main function to simulate Vending Machine functionality
func main() {
	// Initialize items
	item1 := &Item{ID: 1, Name: "Soda", Price: 1.50, Quantity: 10}
	item2 := &Item{ID: 2, Name: "Chips", Price: 1.00, Quantity: 5}
	item3 := &Item{ID: 3, Name: "Chocolate", Price: 2.00, Quantity: 7}

	// Initialize inventory
	inventory := &Inventory{
		items: make(map[int]*Item),
	}
	inventory.addItem(item1)
	inventory.addItem(item2)
	inventory.addItem(item3)

	// Initialize vending machine
	vendingMachine := &VendingMachine{
		inventory: inventory,
	}

	// Display items
	fmt.Println("Available Items:")
	for _, item := range vendingMachine.inventory.items {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f, Quantity: %d\n", item.ID, item.Name, item.Price, item.Quantity)
	}

	// Selecting an item
	vendingMachine.selectItem(1)

	// Display price
	display := &Display{}
	display.displayPrice(vendingMachine.currentSelection)

	// Insert money
	vendingMachine.insertMoney(2.00)

	// Dispense item
	vendingMachine.dispenseItem()

	// Process payment with coin
	coin := &Coin{value: 2.00}
	if coin.acceptPayment(1.50) {
		coin.processPayment()
	}

	// Record transaction
	transaction := &Transaction{
		transactionID: 1,
		dateTime:      time.Now(),
		item:          item1,
		amountPaid:    1.50,
		paymentStatus: "Success",
	}
	transaction.recordTransaction()
	transaction.createReceipt()
}
