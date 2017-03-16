package entity

type (
	ServiceID   int64
	ServiceName string
	ServiceHost string
)

type Service struct {
	ID   ServiceID
	Name ServiceName
	Host ServiceHost
}

type ServiceRepository interface {
	GetAll() ([]*Service, error)
}
