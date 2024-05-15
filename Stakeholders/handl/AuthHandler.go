package handl

import (
	"context"
	"fmt"

	"stakeholders_service/model"
	"stakeholders_service/proto/auth"

	"gorm.io/gorm"
)

type AuthHandler struct {
	auth.UnimplementedStakeholderServiceServer
	DatabaseConnection *gorm.DB
}

func (h AuthHandler) Greet(ctx context.Context, request *auth.Request) (*auth.Response, error) {
	var users []model.User
	if err := h.DatabaseConnection.Table(`stakeholders."Users"`).Find(&users).Error; err != nil {
		return nil, err
	}
	fmt.Println(users)
	return &auth.Response{
		Greeting: fmt.Sprintf("Hi %s!", request.Name),
	}, nil
}
