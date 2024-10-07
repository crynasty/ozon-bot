package product

var allProducts = []Product{ // имитация того что вернет БД
	{Title: "one"},
	{Title: "two"},
	{Title: "three"},
	{Title: "four"},
	{Title: "five"},
}

type Product struct {
	Title string
}