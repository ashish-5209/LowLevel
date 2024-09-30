package main

import "fmt"

// VendingMachine represents the vending machine
type VendingMachine struct {
	balance          float64
	totalSales       float64
	currentSelection *Item
	inventory        *Inventory
}

func (vm *VendingMachine) selectItem(itemID int) {
	item := vm.inventory.getItemDetails(itemID)
	if item != nil && item.Quantity > 0 {
		vm.currentSelection = item
		fmt.Println("Item selected:", item.Name)
	} else {
		fmt.Println("Item not available.")
	}
}

func (vm *VendingMachine) insertMoney(amount float64) {
	vm.balance += amount
	fmt.Printf("Inserted money: %.2f, Current Balance: %.2f\n", amount, vm.balance)
}

func (vm *VendingMachine) dispenseItem() {
	if vm.currentSelection == nil {
		fmt.Println("No item selected.")
		return
	}

	if vm.balance < vm.currentSelection.Price {
		fmt.Println("Insufficient balance.")
		return
	}

	// Dispense item and reduce quantity
	vm.currentSelection.Quantity--
	vm.totalSales += vm.currentSelection.Price
	vm.balance -= vm.currentSelection.Price

	fmt.Println("Dispensing item:", vm.currentSelection.Name)
	vm.currentSelection = nil // Reset selection
}

func (vm *VendingMachine) refund() {
	fmt.Printf("Refunding: %.2f\n", vm.balance)
	vm.balance = 0.0
}

func (vm *VendingMachine) cancel() {
	fmt.Println("Transaction cancelled.")
	vm.refund()
	vm.currentSelection = nil
}
