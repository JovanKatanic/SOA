package handl

import (
	"context"
	"fmt"
	"user_management_service/model"
	"user_management_service/proto/followings"

	"gorm.io/gorm"
)

type FollowingsHandler struct {
	followings.UnimplementedFollowerServiceServer
	DatabaseConnection *gorm.DB
}

func (h FollowingsHandler) GetFollowings(ctx context.Context, request *followings.GetFollowRequest) (*followings.GetFollowResponse, error) {
	var people []model.PeoplePostgre
	fmt.Print("usao, ", request.Id)
	fmt.Println("Entered handler")

	if err := h.DatabaseConnection.Table(`stakeholders."People"`).Find(&people).Error; err != nil {
		return nil, err
	}

	var pbPeople []*followings.People
	for _, p := range people {
		pbPeople = append(pbPeople, &followings.People{
			Ud:         int64(p.ID),
			UserId:     int64(p.UserID),
			Name:       p.Name,
			Surname:    p.Surname,
			Email:      p.Email,
			ProfilePic: p.ProfilePic,
			Biography:  p.Biography,
			Motto:      p.Motto,
			Latitude:   p.Latitude,
			Longitude:  p.Longitude,
		})
	}

	response := &followings.GetFollowResponse{
		People: pbPeople,
	}

	return response, nil
}
