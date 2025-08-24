package rocket

// rocket should contain the defination of our rocket
type Rocket struct {
	ID     string
	Name   string
	Type   string
	Flight int
}

type Store interface{
	
}

// service responsible for updating the rocket inventory
type Service struct {
}

// New returns a new instance of our rocket service
func New() Service {
	return Service{}
}
