package user

type UserInputPort interface {
	findAllUsers() []User
}
