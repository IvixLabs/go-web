package product

//go:generate mockery --name Repository
type Repository interface {
	CreateProduct(p Product)
	FindAllProducts() []Product
	FindProductById(productId string) Product
	UpdateProduct(ps Product)
	DeleteProduct(id string, userId string) int64
	FindProductsByUserId(userId string) []Product
}
