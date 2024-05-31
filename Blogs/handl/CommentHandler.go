package handl

import (
	"blogs_service/model"
	"blogs_service/proto/blogs"
	"context"
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type CommentHandler struct {
	blogs.UnimplementedBlogServiceServer
	DatabaseConnection *gorm.DB
}

func (h BlogHandler) CreateComment(ctx context.Context, request *blogs.Comment) (*blogs.Comment, error) {
	fmt.Println("usao u create commentara")

	var comment = model.Comment{
		Id:           int(request.Id),
		UserId:       int(request.UserId),
		CreationDate: request.CreationDate.AsTime(),
		Description:  request.Description,
		LastEditDate: request.LastEditDate.AsTime(),
		BlogId:       int(request.BlogId),
	}
	if err := h.DatabaseConnection.Table(`blog."Comments"`).Create(&comment).Error; err != nil {
		return nil, err
	}

	response := &blogs.Comment{
		Id:           int32(comment.Id),
		UserId:       int32(comment.UserId),
		CreationDate: request.CreationDate,
		Description:  comment.Description,
		LastEditDate: request.LastEditDate,
		BlogId:       int32(comment.BlogId),
	}

	return response, nil

}

func (h BlogHandler) GetCommentsByBlogIdAsync(ctx context.Context, request *blogs.GetCommentRequest) (*blogs.ListComment, error) {
	blogId := request.BlogId

	var comments []model.Comment

	if err := h.DatabaseConnection.Table(`blog."Comments"`).Where(`"BlogId" = ?`, blogId).Find(&comments).Error; err != nil {
		return nil, err
	}

	protoComments := make([]*blogs.Comment, len(comments))
	for i, comment := range comments {

		protoComments[i] = &blogs.Comment{
			Id:           int32(comment.Id),
			UserId:       int32(comment.UserId),
			CreationDate: timeToTimestamp(comment.CreationDate),
			Description:  comment.Description,
			LastEditDate: timeToTimestamp(comment.LastEditDate),
			BlogId:       int32(comment.BlogId),
		}
	}

	response := &blogs.ListComment{
		Comments: protoComments,
	}

	return response, nil
}

func timeToTimestamp(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}
