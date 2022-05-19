package Auth

type LoginService interface {
	LoginUser(email string, password string) bool
}

type loginInformation struct {
	email    string
	password string
}

func (info *loginInformation) LoginUser(email string, password string) bool {
	return info.email == email && info.password == password
}