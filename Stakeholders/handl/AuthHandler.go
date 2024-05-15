package handl

import (
	"context"
	"fmt"

	"stakeholders_service/model"
	"stakeholders_service/proto/auth"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthHandler struct {
	auth.UnimplementedStakeholderServiceServer
	DatabaseConnection *gorm.DB
}

func ConvertToString(role int) string {
	switch role {
	case 0:
		return "administrator"
	case 1:
		return "author"
	case 2:
		return "tourist"
	default:
		return "unknown"
	}
}

func (h AuthHandler) Greet(ctx context.Context, request *auth.Request) (*auth.Response, error) {
	var user model.User
	if err := h.DatabaseConnection.Table(`stakeholders."Users"`).Where(`"Users"."Username" = ? and "Users"."IsActive" = true`, "stefanstojanovic").First(&user).Error; err != nil {
		return nil, err
	}

	var person model.Person
	if err := h.DatabaseConnection.Table(`stakeholders."People"`).Where(`"People"."UserId" = ?`, user.ID).First(&person).Error; err != nil {
		return nil, err
	}

	claims := jwt.MapClaims{
		"jti":      uuid.New().String(),
		"id":       user.ID,
		"username": user.Username,
		"personId": person.ID,
		"role":     ConvertToString(user.Role),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("your_secret_key")) // Replace "your_secret_key" with your actual secret key
	if err != nil {
		return nil, err
	}
	authenticationResponse := model.AuthenticationTokensDto{
		ID:          user.ID,
		AccessToken: tokenString,
	}

	fmt.Println(authenticationResponse)

	return &auth.Response{
		Greeting: fmt.Sprintf("Hi %s!", request.Name),
	}, nil
}
