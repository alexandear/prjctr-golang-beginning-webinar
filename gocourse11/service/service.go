package service

type Repository interface {
	User() (string, error)
}

type Logger interface{}

type Broker interface{}

type Service struct {
	repository Repository
	Logger     Logger
	broker     Broker
}

func NewService(
	repository Repository,
	logger Logger,
	broker Broker,
) Service {
	return Service{
		repository: repository,
		Logger:     logger,
		broker:     broker,
	}
}

func (s Service) User() error {
	user, err := s.repository.User()
	if err != nil {
		return err
	}
	_ = user
	_ = s.Logger
	_ = s.broker
	return nil
}
