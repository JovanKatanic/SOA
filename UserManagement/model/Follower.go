package model

type Follower struct {
	ID           int                  `json:"id" bson:"id"`
	FollowerId   int                  `json:"followerId" bson:"followerId"`
	FollowedId   int                  `json:"followedId" bson:"followedId"`
	Notification FollowerNotification `json:"notification" bson:"notification"`
}
