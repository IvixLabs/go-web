package property

type UpdateProperty struct {
	Name   string
	IsName bool
}

type Repository interface {
	Create(p Property)
	FindAll() []Property
	FindById(id string) Property
	Update(p Property, updateProperty *UpdateProperty)
	Delete(id string, userId string) int64
	FindByUserId(userId string) []Property
}
