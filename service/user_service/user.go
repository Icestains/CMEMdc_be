package user_service

import "CMEMdc_be/models"

type Auth struct {
	Username string
	Password string
}

func (a *Auth) CheckAuth() (bool, error) {
	return models.CheckAuth(a.Username, a.Password)
}