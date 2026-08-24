package main

import (
	"fmt"
)

type User struct {
	id    int
	name  string
	email string
}

type Credentials struct {
	email    string
	password string
	token    string
}

type AuthProvider interface {
	authenticate(Credentials) error
	name() string
}

type GoogleAuth struct {
}

func (google *GoogleAuth) name() string {
	return "google"
}

func (google *GoogleAuth) authenticate(creds Credentials) error {
	//run some simulation
	return nil
}

type GithubAuth struct {
}

func (github *GithubAuth) authenticate(creds Credentials) error {
	//run some simulation
	return nil
}

func (github *GithubAuth) name() string {
	return "github"
}

type EmailAuth struct{}

func (emailauth *EmailAuth) authenticate(creds Credentials) error {
	return nil
}

func (emailauth *EmailAuth) name() string {
	return "email"
}

type AuthService struct {
	authProvider AuthProvider
}

func (a *AuthService) Login(creds Credentials) error {
	fmt.Printf("login via %s ", a.authProvider.name())
	return a.authProvider.authenticate(creds)
}

func main() {
	fmt.Println("Hello, World!")
	a := AuthService{

		authProvider: &GoogleAuth{},
	}
	creds := Credentials{
		email:    "",
		password: "",
		token:    "",
	}
	a.Login(creds)
}
