package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"solution/model"
	"solution/proto/serviceBook"
	"solution/shared"

	"github.com/go-chi/chi"
	"google.golang.org/grpc"
)

type BookServiceHandler struct {
}

func (BookServiceHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var dto model.BookParams

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		shared.HttpResponseError(w, r, err, http.StatusBadRequest)
		return
	}

	bookProto := &bookpb.CreateBookReq{
		Book: &bookpb.Book{
			Author:      dto.Author,
			Title:       dto.Title,
			Description: dto.Description,
		},
	}

	conn, err := grpc.Dial(":8082", grpc.WithInsecure())
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	defer conn.Close()

	client := bookpb.NewBookServiceClient(conn)
	res, err := client.CreateBook(context.Background(), bookProto)
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	shared.HttpResponseSuccess(w, r, res)
	return
}

func (BookServiceHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	bookId := chi.URLParam(r, "bookid")

	bookProto := &bookpb.ReadBookReq{Id: bookId}

	conn, err := grpc.Dial(":8082", grpc.WithInsecure())
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	defer conn.Close()

	client := bookpb.NewBookServiceClient(conn)
	res, err := client.ReadBook(context.Background(), bookProto)
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	shared.HttpResponseSuccess(w, r, res)
	return
}

func (BookServiceHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var dto model.UpdateBookParams

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		shared.HttpResponseError(w, r, err, http.StatusBadRequest)
		return
	}

	bookProto := &bookpb.UpdateBookReq{
		Book: &bookpb.Book{
			Id:          dto.ID,
			Author:      dto.Author,
			Title:       dto.Title,
			Description: dto.Description,
		},
	}

	conn, err := grpc.Dial(":8082", grpc.WithInsecure())
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	defer conn.Close()

	client := bookpb.NewBookServiceClient(conn)
	res, err := client.UpdateBook(context.Background(), bookProto)
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	shared.HttpResponseSuccess(w, r, res)
	return
}

func (BookServiceHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId := chi.URLParam(r, "bookid")

	bookProto := &bookpb.DeleteBookReq{Id: bookId}

	conn, err := grpc.Dial(":8082", grpc.WithInsecure())
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	defer conn.Close()

	client := bookpb.NewBookServiceClient(conn)
	res, err := client.DeleteBook(context.Background(), bookProto)
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	shared.HttpResponseSuccess(w, r, res)
	return
}