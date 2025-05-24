package model

type InvoiceItem struct {
	Id              int
	Item            string
	Description     string
	Quantity        int
	Price           int
	DiscountedPrice int
	Total           int
}
