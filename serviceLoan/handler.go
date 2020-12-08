package serviceLoan

import (
	"context"
	"fmt"

	"solution/infrastructure"
	"solution/model"
	"solution/proto/serviceLoan"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LoanServiceServer struct {
}

func (s *LoanServiceServer) ReadLoan(ctx context.Context, req *loanpb.ReadLoanReq) (*loanpb.ReadLoanRes, error) {
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	result := infrastructure.Mongodb.Collection("loan").FindOne(ctx, bson.M{"_id": oid})

	data := model.LoanItem{}

	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find loan with Object Id %s: %v", req.GetId(), err))
	}

	response := &loanpb.ReadLoanRes{
		Loan: &loanpb.Loan{
			Id:           oid.Hex(),
			Name:         data.Name,
			BookName:     data.BookName,
			LoanDuration: data.LoanDuration,
		},
	}
	return response, nil
}

func (s *LoanServiceServer) CreateLoan(ctx context.Context, req *loanpb.CreateLoanReq) (*loanpb.CreateLoanRes, error) {
	loan := req.GetLoan()
	data := model.LoanItem{
		Name:         loan.GetName(),
		BookName:     loan.GetBookName(),
		LoanDuration: loan.GetLoanDuration(),
	}

	result, err := infrastructure.Mongodb.Collection("loan").InsertOne(context.Background(), data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	oid := result.InsertedID.(primitive.ObjectID)
	loan.Id = oid.Hex()
	return &loanpb.CreateLoanRes{Loan: loan}, nil
}

func (s *LoanServiceServer) UpdateLoan(ctx context.Context, req *loanpb.UpdateLoanReq) (*loanpb.UpdateLoanRes, error) {
	loan := req.GetLoan()

	oid, err := primitive.ObjectIDFromHex(loan.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Could not convert the supplied loan id to a MongoDB ObjectId: %v", err),
		)
	}

	update := bson.M{
		"name":         loan.GetName(),
		"bookname":     loan.GetBookName(),
		"loanduration": loan.GetLoanDuration(),
	}

	filter := bson.M{"_id": oid}
	option := options.FindOneAndUpdate().SetReturnDocument(1)
	result := infrastructure.Mongodb.Collection("loan").FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, option)

	decoded := model.LoanItem{}
	err = result.Decode(&decoded)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Could not find loan with supplied ID: %v", err),
		)
	}
	return &loanpb.UpdateLoanRes{
		Loan: &loanpb.Loan{
			Id:           decoded.ID.Hex(),
			Name:         decoded.Name,
			BookName:     decoded.BookName,
			LoanDuration: decoded.LoanDuration,
		},
	}, nil
}

func (s *LoanServiceServer) DeleteLoan(ctx context.Context, req *loanpb.DeleteLoanReq) (*loanpb.DeleteLoanRes, error) {
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	_, err = infrastructure.Mongodb.Collection("loan").DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find/delete loan with id %s: %v", req.GetId(), err))
	}
	return &loanpb.DeleteLoanRes{
		Success: true,
	}, nil
}
