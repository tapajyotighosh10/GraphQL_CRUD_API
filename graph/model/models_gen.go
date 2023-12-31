// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type DeleteUser struct {
	UserID string `json:"userId"`
}

type NewUser struct {
	UserID    *string `json:"userId,omitempty"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Dob       string  `json:"dob"`
}

type UpdateUser struct {
	UserID    *string `json:"userId,omitempty"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Dob       string  `json:"dob"`
}

type User struct {
	Error     bool    `json:"error"`
	Message   string  `json:"message"`
	ID        *string `json:"id,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	Dob       *string `json:"dob,omitempty"`
}
