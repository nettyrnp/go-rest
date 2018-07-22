package models

type Identity interface {
	GetID() string
	GetName() string
	GetRole() string
}

type AuthUser struct {
	ID   string
	Name string
	Role string
}

func (au AuthUser) GetID() string {
	return au.ID
}

func (au AuthUser) GetName() string {
	return au.Name
}

func (au AuthUser) GetRole() string {
	return au.Role
}
