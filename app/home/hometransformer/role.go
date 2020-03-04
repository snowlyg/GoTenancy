package hometransformer

type Role struct {
	Id          int
	Name        string
	DisplayName string
	Description string
	Perms       []*Permission
	CreatedAt   string
}
