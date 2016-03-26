package models

type User struct{
	Id int
	Name string
}

type UserSignUpResponse struct{
	Message string
}

type UserFetchResponse struct{
	Message string
	User_details []User
}