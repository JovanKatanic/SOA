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
	Key                string
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

func generateAccessToken(user model.User, person model.Person, key string) (string, error) {
	claims := jwt.MapClaims{
		"jti":      uuid.New().String(),
		"id":       user.ID,
		"username": user.Username,
		"personId": person.ID,
		"role":     ConvertToString(user.Role),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (h AuthHandler) LogIn(ctx context.Context, request *auth.RequestLogIn) (*auth.ResponseLogIn, error) {
	var user model.User
	if err := h.DatabaseConnection.Table(`stakeholders."Users"`).Where(`"Users"."Username" = ? and "Users"."IsActive" = true`, request.Username).First(&user).Error; err != nil {
		return nil, err
	}

	var person model.Person
	if err := h.DatabaseConnection.Table(`stakeholders."People"`).Where(`"People"."UserId" = ?`, user.ID).First(&person).Error; err != nil {
		return nil, err
	}

	tokenString, _ := generateAccessToken(user, person, h.Key)

	return &auth.ResponseLogIn{
		Id:          user.ID,
		AccessToken: tokenString,
	}, nil
}

func (h AuthHandler) RegisterTourist(ctx context.Context, request *auth.RequestRegister) (*auth.ResponseLogIn, error) {

	var user model.User = model.User{
		Username: request.Username,
		Password: request.Password,
		Role:     2,
		IsActive: false,
	}

	dbResult := h.DatabaseConnection.Table(`stakeholders."Users"`).Create(&user)

	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	var person model.Person = model.Person{
		UserID:  user.ID,
		Name:    request.Name,
		Surname: request.Surname,
		Email:   request.Email,
	}
	dbResultPerson := h.DatabaseConnection.Table(`stakeholders."People"`).Create(&person)

	if dbResultPerson.Error != nil {
		return nil, dbResultPerson.Error
	}
	tokenString, _ := generateAccessToken(user, person, h.Key)

	user.EmailVerificationToken = &tokenString
	dbResultToken := h.DatabaseConnection.Table(`stakeholders."Users"`).Save(user)

	if dbResultToken.Error != nil {
		return nil, dbResultToken.Error
	}

	tokenStringAccess, _ := generateAccessToken(user, person, h.Key)

	return &auth.ResponseLogIn{
		Id:          user.ID,
		AccessToken: tokenStringAccess,
	}, nil
}

func (h AuthHandler) ActivateUser(ctx context.Context, request *auth.RequestActivateUser) (*auth.ResponseLogIn, error) {

	token, err := jwt.Parse(request.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.Key), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	id, ok := claims["id"].(int64)
	if !ok {
		return nil, fmt.Errorf("invalid id in token")
	}

	var user model.User
	if err := h.DatabaseConnection.Table(`stakeholders."Users"`).First(&user, id).Error; err != nil {
		return nil, err
	}

	if user.EmailVerificationToken == nil || user.EmailVerificationToken != &request.Token {
		return nil, fmt.Errorf("invalid token")
	}

	expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)

	if time.Now().After(expirationTime) {
		return nil, fmt.Errorf("token has expired")
	}

	user.IsActive = true
	user.EmailVerificationToken = nil

	var person model.Person
	if err := h.DatabaseConnection.Table(`stakeholders."People"`).Where(`"People"."UserId" = ?`, user.ID).First(&person).Error; err != nil {
		return nil, err
	}

	tokenStringAccess, _ := generateAccessToken(user, person, h.Key)

	return &auth.ResponseLogIn{
		Id:          user.ID,
		AccessToken: tokenStringAccess,
	}, nil
}

func (h AuthHandler) ChangePassword(ctx context.Context, request *auth.RequestChangePassword) (*auth.RequestActivateUser, error) {
	return &auth.RequestActivateUser{
		Token: "ahha",
	}, nil
}

func (h AuthHandler) ChangePasswordRequest(ctx context.Context, request *auth.RequestChangePasswordRequest) (*auth.RequestActivateUser, error) {
	return &auth.RequestActivateUser{
		Token: "ahha",
	}, nil
}
