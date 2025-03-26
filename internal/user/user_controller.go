package user

type UserController struct {
	service *UserService
}

func NewUserController(service *UserService) *UserController {
	return &UserController{
		service: service,
	}
}
