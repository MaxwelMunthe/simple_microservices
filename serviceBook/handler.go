package serviceBook

import (
	"context"
	"fmt"

	"solution/infrastructure"
	"solution/model"
	"solution/proto/serviceBook"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BookServiceServer struct {
}

func (s *BookServiceServer) ReadBook(ctx context.Context, req *bookpb.ReadBookReq) (*bookpb.ReadBookRes, error) {
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	result := infrastructure.Mongodb.Collection("book").FindOne(ctx, bson.M{"_id": oid})

	data := model.BookItem{}

	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find book with Object Id %s: %v", req.GetId(), err))
	}

	response := &bookpb.ReadBookRes{
		Book: &bookpb.Book{
			Id:          oid.Hex(),
			Author:      data.Author,
			Title:       data.Title,
			Description: data.Description,
		},
	}
	return response, nil
}

func (s *BookServiceServer) CreateBook(ctx context.Context, req *bookpb.CreateBookReq) (*bookpb.CreateBookRes, error) {
	book := req.GetBook()
	data := model.BookItem{
		Author:      book.GetAuthor(),
		Title:       book.GetTitle(),
		Description: book.GetDescription(),
	}

	result, err := infrastructure.Mongodb.Collection("book").InsertOne(context.Background(), data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	oid := result.InsertedID.(primitive.ObjectID)
	book.Id = oid.Hex()
	return &bookpb.CreateBookRes{Book: book}, nil
}

func (s *BookServiceServer) UpdateBook(ctx context.Context, req *bookpb.UpdateBookReq) (*bookpb.UpdateBookRes, error) {
	book := req.GetBook()

	oid, err := primitive.ObjectIDFromHex(book.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Could not convert the supplied book id to a MongoDB ObjectId: %v", err),
		)
	}

	update := bson.M{
		"author":      book.GetAuthor(),
		"title":       book.GetTitle(),
		"description": book.GetDescription(),
	}

	filter := bson.M{"_id": oid}
	option := options.FindOneAndUpdate().SetReturnDocument(1)
	result := infrastructure.Mongodb.Collection("book").FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, option)

	decoded := model.BookItem{}
	err = result.Decode(&decoded)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Could not find book with supplied ID: %v", err),
		)
	}
	return &bookpb.UpdateBookRes{
		Book: &bookpb.Book{
			Id:          decoded.ID.Hex(),
			Author:      decoded.Author,
			Title:       decoded.Title,
			Description: decoded.Description,
		},
	}, nil
}

func (s *BookServiceServer) DeleteBook(ctx context.Context, req *bookpb.DeleteBookReq) (*bookpb.DeleteBookRes, error) {
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	_, err = infrastructure.Mongodb.Collection("book").DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find/delete book with id %s: %v", req.GetId(), err))
	}
	return &bookpb.DeleteBookRes{
		Success: true,
	}, nil
}
