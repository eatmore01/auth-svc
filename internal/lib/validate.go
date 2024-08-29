package lib

import (
	"errors"

	emailverifier "github.com/AfterShip/email-verifier"
)

var (
	InvalidSyntaxEmail = errors.New("Invalid syntax Email")
	VerifyEmailError   = errors.New("Verify Email Error")
	PasswordIsEmpty    = errors.New("Password is empty")
	IdIsEmpty          = errors.New("Id is empty")
	EmailIsEmpty       = errors.New("Email is empty")
	UserNameIsEmpty    = errors.New("User name is empty")
)

var (
	verifier = emailverifier.NewVerifier().DisableCatchAllCheck()
)

func ValidateLoginRequest(r LoginRequest) error {
	if r.Email == "" {
		return EmailIsEmpty
	}

	res, err := verifier.Verify(r.Email)
	if !res.Syntax.Valid {
		return InvalidSyntaxEmail
	}
	if err != nil {
		return VerifyEmailError
	}

	if r.Password == "" {
		return PasswordIsEmpty
	}
	if r.Id == "" {
		return IdIsEmpty
	}

	return nil
}

func ValidateRegisterRequest(r RegisterRequest) error {
	if r.Email == "" {
		return EmailIsEmpty
	}

	res, err := verifier.Verify(r.Email)
	if !res.Syntax.Valid {
		return InvalidSyntaxEmail
	}
	if err != nil {
		return VerifyEmailError
	}

	if r.Password == "" {
		return PasswordIsEmpty
	}
	if r.User_name == "" {
		return UserNameIsEmpty
	}

	return nil
}
