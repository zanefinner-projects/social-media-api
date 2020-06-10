package users

//UserCredentials is the struct to organize the user json when getting tokens
type UserCredentials struct {
	Username string
	Password string
}

//UserDataForUserCreate is the struct to organize the user json when creating an account
type UserDataForUserCreate struct {
	Username string
	Password string
	Email    string
	//more to come...
}
