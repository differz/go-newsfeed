package repository

import (
	"github.com/VitaliiHurin/go-newsfeed/entity"
	"upper.io/db.v3/lib/sqlbuilder"
)

type serviceTable struct {
	ID   int64  `db:"id,omitempty"`
	Name string `db:"name"`
	Host string `db:"host"`
}

func assembleService(t *serviceTable) *entity.Service {
	return &entity.Service{
		ID:   entity.ServiceID(t.ID),
		Name: entity.ServiceName(t.Name),
		Host: entity.ServiceHost(t.Host),
	}
}

func newServiceTable(r *entity.Service) *serviceTable {
	return &serviceTable{
		ID:   int64(r.ID),
		Name: string(r.Name),
		Host: string(r.Host),
	}
}

type serviceRepository struct {
	DB sqlbuilder.Database
}

func NewServiceRepository(DB sqlbuilder.Database) entity.ServiceRepository {
	return &serviceRepository{
		DB: DB,
	}
}

func (r *serviceRepository) GetAll() ([]*entity.Service, error) {
	var s []*serviceTable
	res := r.DB.Collection("service").Find()
	err := res.All(&s)
	if err != nil {
		return nil, err
	}
	var services []*entity.Service
	for _, v := range s {
		services = append(services, assembleService(v))
	}

	return services, nil
}
