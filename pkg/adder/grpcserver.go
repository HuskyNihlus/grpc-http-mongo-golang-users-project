package adder

import (
	"PROJECT/pkg/api"
	"PROJECT/pkg/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//GRPCServer ...
type GRPCServer struct {
	api.UnimplementedAdderServer

	Collection *mongo.Collection
}

func New(
	collection *mongo.Collection,
) *GRPCServer {
	return &GRPCServer{
		Collection: collection,
	}
}

// Add ...
func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	fmt.Println("Збс работает")
	res, err := s.Collection.InsertOne(ctx, model.NewUserFromAddRequest(req))
	if err != nil {
		return nil, err
	}
	return &api.AddResponse{Id: res.InsertedID.(primitive.ObjectID).String()}, nil
}


func (s *GRPCServer) GetAll(ctx context.Context, request *api.GetAllRequest) (*api.GetAllResponse, error) {
	res, err := s.Collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	users := make([]*api.User, 0)
	for res.Next(ctx) {
		var user model.User
		err = res.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user.ToProtoUser())
	}

	return &api.GetAllResponse{Users: users}, nil
}

func (s *GRPCServer) GetByID(ctx context.Context, request *api.GetByIDRequest) (*api.GetByIDResponse, error) {
	panic("implement me")
}