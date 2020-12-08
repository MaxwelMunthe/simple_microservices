package serviceAccount

import (
	"context"
	"log"
	"net"
	"net/url"
	"testing"

	"solution/proto/serviceAccount"
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

	accountpb.RegisterAccountServiceServer(server, &AccountServiceServer{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestAccountServiceServer_CreateAccount(t *testing.T) {
	expectedResult := &accountpb.CreateAccountRes{Account: &accountpb.Account{
		Name:        "test tost",
		Address:     "test street",
		PhoneNumber: "+62889900",
	}}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := accountpb.NewAccountServiceClient(conn)
	request := &accountpb.CreateAccountReq{Account: &accountpb.Account{
		Name:        "test tost",
		Address:     "test street",
		PhoneNumber: "+62889900",
	}}

	response, err := client.CreateAccount(ctx, request)
	assert.Equal(t, expectedResult.Account.Name, response.Account.Name)
	assert.Equal(t, expectedResult.Account.Address, response.Account.Address)
	assert.Equal(t, expectedResult.Account.PhoneNumber, response.Account.PhoneNumber)
	assert.NoError(t, err)
}

func TestAccountServiceServer_ReadAccount_NotFoundData(t *testing.T) {
	expectedResult := "rpc error: code = NotFound desc = Could not find account with Object Id 5f81e55d9f9ff3494f7a1b10: mongo: no documents in result"

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := accountpb.NewAccountServiceClient(conn)
	request := &accountpb.ReadAccountReq{Id: "5f81e55d9f9ff3494f7a1b10"}

	_, err = client.ReadAccount(ctx, request)
	assert.Equal(t, expectedResult, err.Error())
}

func TestAccountServiceServer_UpdateAccount_NotFoundData(t *testing.T) {
	expectedResult := "rpc error: code = NotFound desc = Could not find account with supplied ID: mongo: no documents in result"

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := accountpb.NewAccountServiceClient(conn)
	request := &accountpb.UpdateAccountReq{Account: &accountpb.Account{
		Id:          "5f81e55d9f9ff3494f7a1b10",
		Name:        "test",
		Address:     "test street",
		PhoneNumber: "+62889900",
	}}

	_, err = client.UpdateAccount(ctx, request)
	assert.Equal(t, expectedResult, err.Error())
}

func TestAccountServiceServer_DeleteAccount_NotFoundData(t *testing.T) {
	expectedResult := true
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := accountpb.NewAccountServiceClient(conn)
	request := &accountpb.DeleteAccountReq{Id:"5f81e55d9f9ff3494f7a1b10"}

	response, err := client.DeleteAccount(ctx, request)
	assert.Equal(t, expectedResult, response.Success)
	assert.NoError(t, err)
}
