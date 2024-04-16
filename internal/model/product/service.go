package product

type Service interface {
	FindAll() []Product
	FindByUserId(userId string) []Product
	FindById(productId string) Product
}

type service struct {
	productRepository Repository
}

func NewService(productRepository Repository) Service {
	return &service{productRepository: productRepository}
}

func (service *service) FindAll() []Product {
	return service.productRepository.FindAllProducts()
}

func (service *service) FindByUserId(userId string) []Product {
	return service.productRepository.FindProductsByUserId(userId)
}

func (service *service) FindById(productId string) Product {
	return service.productRepository.FindProductById(productId)
}
