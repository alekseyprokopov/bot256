package product

var allProducts = []Product{
	{Name: "one"},
	{Name: "two"},
	{Name: "three"},
	{Name: "four"},
	{Name: "five"},
}


type Product struct {
	Name string
}