package routes

import (
	"solution/handler"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func RegisterRoutes(r *chi.Mux) *chi.Mux {
	accountGrpc := handler.AccountServiceHandler{}
	bookGrpc := handler.BookServiceHandler{}
	loanGrpc := handler.LoanServiceHandler{}

	r.Use(middleware.Logger)
	r.Route("/account", func(r chi.Router) {
		r.Post("/", accountGrpc.CreateAccount)
		r.Put("/", accountGrpc.UpdateAccount)
		r.Get("/{accountid}", accountGrpc.GetAccount)
		r.Post("/delete/{accountid}", accountGrpc.DeleteAccount)
	})

	r.Route("/book", func(r chi.Router) {
		r.Post("/", bookGrpc.CreateBook)
		r.Put("/", bookGrpc.UpdateBook)
		r.Get("/{bookid}", bookGrpc.GetBook)
		r.Post("/delete/{bookid}", bookGrpc.DeleteBook)
	})

	r.Route("/loan", func(r chi.Router) {
		r.Post("/", loanGrpc.CreateLoan)
		r.Put("/", loanGrpc.UpdateLoan)
		r.Get("/{loanid}", loanGrpc.GetLoan)
		r.Post("/delete/{loanid}", loanGrpc.DeleteLoan)
	})
	return r
}
