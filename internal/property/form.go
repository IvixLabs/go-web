package property

type CreateForm struct {
	Name string `validate:"required" json:"name"`
}
