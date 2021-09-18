package model

import (
	"PROJECT/pkg/api"
	"time"
)

type User struct {
	Id            string    `bson:"_id,omitempty"`
	Name          string    `bson:"name,omitempty"`
	LastName      string    `bson:"lastName,omitempty"`
	Age           int32     `bson:"age,omitempty"`
	RecordingDate time.Time `bson:"recordingDate,omitempty"`
}

func NewUserFromAddRequest(req *api.AddRequest) *User {
	return &User{
		Name:          req.Name,
		LastName:      req.LastName,
		Age:           req.Age,
		RecordingDate: time.Now(),
	}
}

func (u User) ToProtoUser() *api.User {
	return &api.User{
		Id:            u.Id,
		Name:          u.Name,
		LastName:      u.LastName,
		Age:           u.Age,
		RecordingDate: u.RecordingDate.Unix(),
	}
}