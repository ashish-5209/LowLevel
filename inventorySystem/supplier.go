package main

import "fmt"

// Supplier struct representing a supplier
type Supplier struct {
	SupplierID int
	Name       string
	Contact    string
	Address    string
}

// Methods for Supplier
func (s *Supplier) AddSupplier() {
	fmt.Printf("Supplier %s added successfully.\n", s.Name)
}

func (s *Supplier) UpdateSupplier(name string, contact string, address string) {
	s.Name = name
	s.Contact = contact
	s.Address = address
	fmt.Printf("Supplier %d updated successfully.\n", s.SupplierID)
}

func (s *Supplier) DeleteSupplier() {
	fmt.Printf("Supplier %d deleted successfully.\n", s.SupplierID)
}

func (s *Supplier) GetSupplierDetails() {
	fmt.Printf("Supplier ID: %d, Name: %s, Contact: %s, Address: %s\n", s.SupplierID, s.Name, s.Contact, s.Address)
}
