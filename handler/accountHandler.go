package handler

import (
	"encoding/json"
	"net/http"

	"solution/model"
	"solution/proto/serviceAccount"
	"solution/shared"

	"github.com/go-chi/chi"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type AccountServiceHandler struct {
}

// Create Account
func (AccountServiceHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var dto model.AccountParams

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		shared.HttpResponseError(w, r, err, http.StatusBadRequest)
		return
	}

	accountProto := &accountpb.CreateAccountReq{
		Account: &accountpb.Account{
			Name:        dto.Name,
			Address:     dto.Address,
			PhoneNumber: dto.PhoneNumber,
		},
	}

	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	defer conn.Close()

	client := accountpb.NewAccountServiceClient(conn)
	res, err := client.CreateAccount(context.Background(), accountProto)
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	shared.HttpResponseSuccess(w, r, res)
	return
}

// Get Account
func (AccountServiceHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	accountId := chi.URLParam(r, "accountid")

	accountProto := &accountpb.ReadAccountReq{Id: accountId}

	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	defer conn.Close()

	client := accountpb.NewAccountServiceClient(conn)
	res, err := client.ReadAccount(context.Background(), accountProto)
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	shared.HttpResponseSuccess(w, r, res)
	return
}

func (AccountServiceHandler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	var dto model.UpdateAccountParams

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		shared.HttpResponseError(w, r, err, http.StatusBadRequest)
		return
	}

	accountProto := &accountpb.UpdateAccountReq{
		Account: &accountpb.Account{
			Id:          dto.ID,
			Name:        dto.Name,
			Address:     dto.Address,
			PhoneNumber: dto.PhoneNumber,
		},
	}

	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	defer conn.Close()

	client := accountpb.NewAccountServiceClient(conn)
	res, err := client.UpdateAccount(context.Background(), accountProto)
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	shared.HttpResponseSuccess(w, r, res)
	return
}

func (AccountServiceHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	accountId := chi.URLParam(r, "accountid")

	accountProto := &accountpb.DeleteAccountReq{Id: accountId}

	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	defer conn.Close()

	client := accountpb.NewAccountServiceClient(conn)
	res, err := client.DeleteAccount(context.Background(), accountProto)
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	shared.HttpResponseSuccess(w, r, res)
	return
}