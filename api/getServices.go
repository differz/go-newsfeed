package api

import "github.com/VitaliiHurin/go-newsfeed/entity"

func (a *API) GetServices() ([]*entity.Service, error) {
	return a.services.GetAll()
}