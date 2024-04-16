package product

type Repository interface {
	Create(p Product)
	FindAllProducts() []Product
	FindProductById(productId string) Product
	UpdateProduct(p Product, updateProduct *UpdateProductArg)
	DeleteProduct(id string, userId string) int64
	FindProductsByUserId(userId string) []Product
}

type UpdateProductArg struct {
	Title   string
	IsTitle bool
	Price   int
	IsPrice bool
	Brand   string
	IsBrand bool
}
