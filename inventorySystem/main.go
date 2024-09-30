// +--------------------+
// |      Product       |
// |--------------------|
// | - productID: int   |
// | - name: string     |
// | - description: string |
// | - price: float     |
// | - quantity: int    |
// | - supplierID: int  |
// |--------------------|
// | + addProduct()     |
// | + updateProduct()  |
// | + deleteProduct()  |
// | + getProductDetails() |
// +--------------------+
//            |
//            |
//            v
// +--------------------+
// |     Supplier       |
// |--------------------|
// | - supplierID: int  |
// | - name: string     |
// | - contact: string   |
// | - address: string  |
// |--------------------|
// | + addSupplier()    |
// | + updateSupplier() |
// | + deleteSupplier() |
// | + getSupplierDetails() |
// +--------------------+
//            |
//            |
//            v
// +--------------------+
// |      Category      |
// |--------------------|
// | - categoryID: int  |
// | - name: string     |
// | - description: string |
// |--------------------|
// | + addCategory()    |
// | + updateCategory() |
// | + deleteCategory() |
// | + getCategoryDetails() |
// +--------------------+
//            |
//            |
//            v
// +--------------------+
// |     Inventory      |
// |--------------------|
// | - inventoryID: int |
// | - productID: int   |
// | - quantity: int     |
// | - location: string  |
// |--------------------|
// | + addInventory()   |
// | + updateInventory() |
// | + deleteInventory() |
// | + getInventoryDetails() |
// +--------------------+
//            |
//            |
//            v
// +--------------------+
// |    Transaction      |
// |---------------------|
// | - transactionID: int|
// | - productID: int    |
// | - quantity: int      |
// | - transactionType: string | // e.g., "In", "Out"
// | - transactionDate: time.Time |
// |---------------------|
// | + addTransaction()  |
// | + getTransactionDetails() |
// +---------------------+

package main

import "time"

// Sample Usage
func main() {
	// Create a supplier
	supplier := Supplier{SupplierID: 1, Name: "ABC Supplies", Contact: "1234567890", Address: "123 Main St"}
	supplier.AddSupplier()

	// Create a product
	product := Product{ProductID: 1, Name: "Widget", Description: "A useful widget", Price: 19.99, Quantity: 100, SupplierID: supplier.SupplierID}
	product.AddProduct()
	product.GetProductDetails()

	// Update product
	product.UpdateProduct("Super Widget", "An even more useful widget", 24.99, 80)
	product.GetProductDetails()

	// Create an inventory record
	inventory := Inventory{InventoryID: 1, ProductID: product.ProductID, Quantity: 50, Location: "Warehouse A"}
	inventory.AddInventory()
	inventory.GetInventoryDetails()

	// Create a transaction for adding stock
	transaction := Transaction{TransactionID: 1, ProductID: product.ProductID, Quantity: 20, TransactionType: "In", TransactionDate: time.Now()}
	transaction.AddTransaction()
	transaction.GetTransactionDetails()

	// Create a transaction for removing stock
	transactionOut := Transaction{TransactionID: 2, ProductID: product.ProductID, Quantity: 10, TransactionType: "Out", TransactionDate: time.Now()}
	transactionOut.AddTransaction()
	transactionOut.GetTransactionDetails()

	// Delete product and supplier
	product.DeleteProduct()
	supplier.DeleteSupplier()
}
