package serviceLoan

import (
	"context"
	"log"
	"net"
	"net/url"
	"testing"

	"solution/proto/serviceLoan"
	"solution/infrastructure"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func init() {
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	env := infrastructure.Environment{}
	env.SetEnvironment()
	env.LoadConfig()
	env.InitMongoDB()
}

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	loanpb.RegisterLoanServiceServer(server, &LoanServiceServer{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestLoanServiceServer_CreateLoan(t *testing.T) {
	expectedResult := &loanpb.CreateLoanRes{Loan: &loanpb.Loan{
		Name:         "test tost",
		BookName:     "test book",
		LoanDuration: "3",
	}}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := loanpb.NewLoanServiceClient(conn)
	request := &loanpb.CreateLoanReq{Loan: &loanpb.Loan{
		Name:         "test tost",
		BookName:     "test book",
		LoanDuration: "3",
	}}

	response, err := client.CreateLoan(ctx, request)
	assert.Equal(t, expectedResult.Loan.Name, response.Loan.Name)
	assert.Equal(t, expectedResult.Loan.BookName, response.Loan.BookName)
	assert.Equal(t, expectedResult.Loan.LoanDuration, response.Loan.LoanDuration)
	assert.NoError(t, err)
}

func TestLoanServiceServer_ReadLoan_NotFoundData(t *testing.T) {
	expectedResult := "rpc error: code = NotFound desc = Could not find loan with Object Id 5f81e55d9f9ff3494f7a1b10: mongo: no documents in result"

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := loanpb.NewLoanServiceClient(conn)
	request := &loanpb.ReadLoanReq{Id: "5f81e55d9f9ff3494f7a1b10"}

	_, err = client.ReadLoan(ctx, request)
	assert.Equal(t, expectedResult, err.Error())
}

func TestLoanServiceServer_UpdateLoan_NotFoundData(t *testing.T) {
	expectedResult := "rpc error: code = NotFound desc = Could not find loan with supplied ID: mongo: no documents in result"

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := loanpb.NewLoanServiceClient(conn)
	request := &loanpb.UpdateLoanReq{Loan: &loanpb.Loan{
		Id:          "5f81e55d9f9ff3494f7a1b10",
		Name:         "test tost",
		BookName:     "test book",
		LoanDuration: "3",
	}}

	_, err = client.UpdateLoan(ctx, request)
	assert.Equal(t, expectedResult, err.Error())
}

func TestLoanServiceServer_DeleteLoan_NotFoundData(t *testing.T) {
	expectedResult := true
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := loanpb.NewLoanServiceClient(conn)
	request := &loanpb.DeleteLoanReq{Id: "5f81e55d9f9ff3494f7a1b10"}

	response, err := client.DeleteLoan(ctx, request)
	assert.Equal(t, expectedResult, response.Success)
	assert.NoError(t, err)
}
