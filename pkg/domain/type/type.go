package _type

type Type struct {
	ID uint `json:"id"`
	Name string `json:"name"`
}

type TypeRepository interface {
	Create(Type) (int, error)
	Read()([]Type,error)
	Update(Type,string) error
	Delete(string) error
}