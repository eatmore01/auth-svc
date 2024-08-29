package tests

import (
	"auth/service/internal/lib"
	"testing"
)

type LR = lib.LoginRequest
type RR = lib.RegisterRequest

type Cases[T any, V any] struct {
	value    T
	expected V
	testName string
}

var LoginCases = []Cases[LR, error]{
	{value: LR{Id: "test", Email: "test@gmail.com", Password: "test"},
		expected: nil,
		testName: "Test valid login request",
	},
	{value: LR{Id: "", Email: "test@gmail.com", Password: "test"},
		expected: lib.IdIsEmpty,
		testName: "test empty id",
	},
	{value: LR{Id: "test", Email: "", Password: "test"},
		expected: lib.EmailIsEmpty,
		testName: "test empty email",
	},
	{value: LR{Id: "test", Email: "test@gmail.com", Password: ""},
		expected: lib.PasswordIsEmpty,
		testName: "test empty password",
	},
	{value: LR{Id: "test", Email: "test@qewqweqweq", Password: "test"},
		expected: lib.InvalidSyntaxEmail,
		testName: "test invalid syntax email",
	},
	{value: LR{Id: "test", Email: "test@qewqweqweq.com", Password: "test"},
		expected: lib.VerifyEmailError,
		testName: "test verify email error",
	},
}

var RegisterCases = []Cases[RR, error]{

	{value: RR{Email: "test@gmail.com", User_name: "test", Password: "test"},
		expected: nil,
		testName: "Test valid register request",
	},
	{value: RR{Email: "", User_name: "test", Password: "test"},
		expected: lib.EmailIsEmpty,
		testName: "test empty email",
	},
	{value: RR{Email: "test@gmail.com", User_name: "", Password: "test"},
		expected: lib.UserNameIsEmpty,
		testName: "test empty usr name",
	},
	{value: RR{Email: "test@gmail.com", User_name: "test", Password: ""},
		expected: lib.PasswordIsEmpty,
		testName: "test empty password",
	},
	{value: RR{Email: "test@qweqweqweom", User_name: "test", Password: ""},
		expected: lib.InvalidSyntaxEmail,
		testName: "test invalid email syntax",
	},
	{value: RR{Email: "test@qweqweqwe.com", User_name: "test", Password: ""},
		expected: lib.VerifyEmailError,
		testName: "test verify email error",
	},
}

func TestValidateLoginFunction(t *testing.T) {
	for _, c := range LoginCases {
		t.Run(c.testName, func(t *testing.T) {
			err := lib.ValidateLoginRequest(c.value)
			if err != c.expected {
				t.Errorf("Error: %v", err)
			}
		})
	}

}

func TestValidateRegisterFunction(t *testing.T) {
	for _, c := range RegisterCases {
		t.Run(c.testName, func(t *testing.T) {
			err := lib.ValidateRegisterRequest(c.value)
			if err != c.expected {
				t.Errorf("Error: %v", err)
			}
		})
	}
}
