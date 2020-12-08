package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"solution/model"
	"solution/proto/serviceLoan"
	"solution/shared"

	"github.com/go-chi/chi"
	"google.golang.org/grpc"
)

type LoanServiceHandler struct {
}

func (LoanServiceHandler) CreateLoan(w http.ResponseWriter, r *http.Request) {
	var dto model.LoanParams

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		shared.HttpResponseError(w, r, err, http.StatusBadRequest)
		return
	}

	loanProto := &loanpb.CreateLoanReq{
		Loan: &loanpb.Loan{
			Name:         dto.Name,
			BookName:     dto.BookName,
			LoanDuration: dto.LoanDuration,
		},
	}

	conn, err := grpc.Dial(":8083", grpc.WithInsecure())
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	defer conn.Close()

	client := loanpb.NewLoanServiceClient(conn)
	res, err := client.CreateLoan(context.Background(), loanProto)
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	shared.HttpResponseSuccess(w, r, res)
	return
}

func (LoanServiceHandler) GetLoan(w http.ResponseWriter, r *http.Request) {
	loanId := chi.URLParam(r, "loanid")

	loanProto := &loanpb.ReadLoanReq{Id: loanId}

	conn, err := grpc.Dial(":8083", grpc.WithInsecure())
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	defer conn.Close()

	client := loanpb.NewLoanServiceClient(conn)
	res, err := client.ReadLoan(context.Background(), loanProto)
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	shared.HttpResponseSuccess(w, r, res)
	return
}

func (LoanServiceHandler) UpdateLoan(w http.ResponseWriter, r *http.Request) {
	var dto model.UpdateLoanParams

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		shared.HttpResponseError(w, r, err, http.StatusBadRequest)
		return
	}

	loanProto := &loanpb.UpdateLoanReq{
		Loan: &loanpb.Loan{
			Id:           dto.ID,
			Name:         dto.Name,
			BookName:     dto.BookName,
			LoanDuration: dto.LoanDuration,
		},
	}

	conn, err := grpc.Dial(":8083", grpc.WithInsecure())
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	defer conn.Close()

	client := loanpb.NewLoanServiceClient(conn)
	res, err := client.UpdateLoan(context.Background(), loanProto)
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	shared.HttpResponseSuccess(w, r, res)
	return
}

func (LoanServiceHandler) DeleteLoan(w http.ResponseWriter, r *http.Request) {
	loanId := chi.URLParam(r, "loanid")

	loanProto := &loanpb.DeleteLoanReq{Id: loanId}

	conn, err := grpc.Dial(":8083", grpc.WithInsecure())
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	defer conn.Close()

	client := loanpb.NewLoanServiceClient(conn)
	res, err := client.DeleteLoan(context.Background(), loanProto)
	if err != nil {
		shared.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	shared.HttpResponseSuccess(w, r, res)
	return
}
