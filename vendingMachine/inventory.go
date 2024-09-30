package main

// Inventory stores items in the vending machine
type Inventory struct {
	items map[int]*Item
}

func (inv *Inventory) addItem(item *Item) {
	inv.items[item.ID] = item
}

func (inv *Inventory) removeItem(itemID int) {
	delete(inv.items, itemID)
}

func (inv *Inventory) getItemDetails(itemID int) *Item {
	return inv.items[itemID]
}

func (inv *Inventory) updateStock(itemID, quantity int) {
	if item, exists := inv.items[itemID]; exists {
		item.updateQuantity(quantity)
	}
}
