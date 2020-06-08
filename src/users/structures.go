package users

type userCredentials struct {
	Username string
	Password string
}

//UserDataForUserCreate is the struct to organize the user json
type UserDataForUserCreate struct {
	Username string
	Password string
	Email    string
	//more to come...
}
