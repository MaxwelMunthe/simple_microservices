package serviceAccount

import (
	"context"
	"fmt"

	"solution/infrastructure"
	"solution/model"
	"solution/proto/serviceAccount"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccountServiceServer struct {
}

func (s *AccountServiceServer) ReadAccount(ctx context.Context, req *accountpb.ReadAccountReq) (*accountpb.ReadAccountRes, error) {
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	result := infrastructure.Mongodb.Collection("account").FindOne(ctx, bson.M{"_id": oid})

	data := model.AccountItem{}

	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find account with Object Id %s: %v", req.GetId(), err))
	}

	response := &accountpb.ReadAccountRes{
		Account: &accountpb.Account{
			Id:          oid.Hex(),
			Name:        data.Name,
			Address:     data.Address,
			PhoneNumber: data.PhoneNumber,
		},
	}
	return response, nil
}

func (s *AccountServiceServer) CreateAccount(ctx context.Context, req *accountpb.CreateAccountReq) (*accountpb.CreateAccountRes, error) {
	account := req.GetAccount()
	data := model.AccountItem{
		Name:        account.GetName(),
		Address:     account.GetAddress(),
		PhoneNumber: account.GetPhoneNumber(),
	}

	result, err := infrastructure.Mongodb.Collection("account").InsertOne(context.Background(), data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	oid := result.InsertedID.(primitive.ObjectID)
	account.Id = oid.Hex()
	return &accountpb.CreateAccountRes{Account: account}, nil
}

func (s *AccountServiceServer) UpdateAccount(ctx context.Context, req *accountpb.UpdateAccountReq) (*accountpb.UpdateAccountRes, error) {
	account := req.GetAccount()

	oid, err := primitive.ObjectIDFromHex(account.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Could not convert the supplied account id to a MongoDB ObjectId: %v", err),
		)
	}

	update := bson.M{
		"name":         account.GetName(),
		"address":      account.GetAddress(),
		"phonenumber": account.GetPhoneNumber(),
	}

	filter := bson.M{"_id": oid}
	option := options.FindOneAndUpdate().SetReturnDocument(1)
	result := infrastructure.Mongodb.Collection("account").FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, option)

	decoded := model.AccountItem{}
	err = result.Decode(&decoded)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Could not find account with supplied ID: %v", err),
		)
	}
	return &accountpb.UpdateAccountRes{
		Account: &accountpb.Account{
			Id:          decoded.ID.Hex(),
			Name:        decoded.Name,
			Address:     decoded.Address,
			PhoneNumber: decoded.PhoneNumber,
		},
	}, nil
}

func (s *AccountServiceServer) DeleteAccount(ctx context.Context, req *accountpb.DeleteAccountReq) (*accountpb.DeleteAccountRes, error) {
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	_, err = infrastructure.Mongodb.Collection("account").DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find/delete account with id %s: %v", req.GetId(), err))
	}
	return &accountpb.DeleteAccountRes{
		Success: true,
	}, nil
}
