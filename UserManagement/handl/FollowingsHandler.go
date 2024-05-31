package handl

import (
	"context"
	"fmt"
	"log"
	"user_management_service/proto/followings"
	"user_management_service/repository"
)

type FollowingsHandler struct {
	followings.UnimplementedFollowerServiceServer
	logger *log.Logger
	repo   *repository.FollowerRepository
}

func NewFollowingsHandler(log *log.Logger, repo *repository.FollowerRepository) *FollowingsHandler {
	return &FollowingsHandler{followings.UnimplementedFollowerServiceServer{}, log, repo}
}

func (h FollowingsHandler) GetFollowings(ctx context.Context, request *followings.GetFollowRequest) (*followings.GetFollowResponse, error) {
	h.logger.Printf("WENT IN!!!!!!")
	fmt.Println("Usao u getFollowings")

	followingss, err := h.repo.GetFollowedPersonsById(int(request.Id))
	if err != nil {
		h.logger.Print("Database exception: ", err)
	}

	fmt.Println(followingss)

	var peopleSlice []*followings.People
	for _, person := range followingss {
		peopleSlice = append(peopleSlice,
			&followings.People{
				Id:         person.ID,
				UserId:     person.UserId,
				Name:       person.Name,
				Surname:    person.Surname,
				Email:      person.Email,
				ProfilePic: person.ProfilePic,
				Biography:  person.Biography,
				Motto:      person.Motto,
				Latitude:   person.Latitude,
				Longitude:  person.Longitude,
			})
	}

	response := &followings.GetFollowResponse{
		People: peopleSlice,
	}

	return response, nil
}
