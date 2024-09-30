package main

import "fmt"

// Category struct representing a product category
type Category struct {
	CategoryID  int
	Name        string
	Description string
}

// Methods for Category
func (c *Category) AddCategory() {
	fmt.Printf("Category %s added successfully.\n", c.Name)
}

func (c *Category) UpdateCategory(name string, description string) {
	c.Name = name
	c.Description = description
	fmt.Printf("Category %d updated successfully.\n", c.CategoryID)
}

func (c *Category) DeleteCategory() {
	fmt.Printf("Category %d deleted successfully.\n", c.CategoryID)
}

func (c *Category) GetCategoryDetails() {
	fmt.Printf("Category ID: %d, Name: %s\n", c.CategoryID, c.Name)
}
