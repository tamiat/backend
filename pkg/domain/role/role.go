package role

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}

type RoleRepository interface {
	Create(Role) (int, error)
	Read() ([]Role, error)
	Delete(int) error
}
