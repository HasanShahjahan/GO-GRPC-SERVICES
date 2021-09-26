package rocket

import "context"

//Rocket - should contain the definition of our
//rocket
type Rocket struct {
	ID      string
	Name    string
	Type    string
	Flights int
}

//Store - defines the interface we expect our database implementation to follow
type Store interface {
	GetRocketByID(id string) (Rocket, error)
	InsertRocket(rocket Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

// Service -our rocket service, responsible for updating the rocket inventory
type Service struct {
	Store Store
}

//New - returns a new instance of our rocket service
func New(store Store) Service {
	return Service{
		Store: store,
	}
}

//GetRocketByID - retrieves a rocket based on the ID from the  store
func (s Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	rocket, err := s.Store.GetRocketByID(id)
	if err != nil {
		return Rocket{}, nil
	}
	return rocket, nil
}

// InsertRocket - inserts a new rocket into the store
func (s Service) InsertRocket(ctx context.Context, rocket Rocket) (Rocket, error) {
	rocket, err := s.Store.InsertRocket(rocket)
	if err != nil {
		return Rocket{}, nil
	}
	return rocket, nil
}

// DeleteRocket - deletes a rocket from our inventory
func (s Service) DeleteRocket(ctx context.Context, id string) error {
	err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}
	return nil
}
