package gholam

type Gholam interface {
	Register(email string, password string) (string, error)
	Login(email string, password string) (string, error)
	Logout(token string) error
}

type gholam struct{}

func New() Gholam {
	return nil
}
