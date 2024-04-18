package user

type UserUseCase struct {
	repository UserOutputPort
}

func (user *UserUseCase) findAllUsers() []User {
	return []User{
		{
			Name: "Lucas",
		},
	}
}

func NewUseCase(repository UserOutputPort) *UserUseCase {
	return &UserUseCase{
		repository: repository,
	}
}
