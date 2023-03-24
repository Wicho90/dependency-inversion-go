package infrastructure

import (
	"github.com/wicho90/dependency-inversion-go/application"
)

var inMemoryUserRepository = NewInMemoryUserRepository()
var userService = application.NewUserService(inMemoryUserRepository)
var UserControllerImpl = NewUserController(*userService)
