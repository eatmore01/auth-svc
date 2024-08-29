package tests

import (
	"auth/service/internal/domain/model"
	"auth/service/internal/lib"
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
)

var mockModel = model.User{
	Id:       uuid.New().String(),
	Email:    "test@t.com",
	UserName: "test",
	Passhash: []byte("test"),
}

const (
	duraction = time.Hour
)

func TestJwtFunction(t *testing.T) {
	token, err := lib.NewJWTToken(mockModel, "Secret", duraction)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	log.Print(token)
}
