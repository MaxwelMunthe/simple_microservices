package serviceBook

import (
	"context"
	"log"
	"net"
	"net/url"
	"testing"

	"solution/proto/serviceBook"

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

	bookpb.RegisterBookServiceServer(server, &BookServiceServer{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestBookServiceServer_CreateBook(t *testing.T) {
	expectedResult := &bookpb.CreateBookRes{Book: &bookpb.Book{
		Author:      "test tost",
		Title:       "test book",
		Description: "this is book",
	}}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := bookpb.NewBookServiceClient(conn)
	request := &bookpb.CreateBookReq{Book: &bookpb.Book{
		Author:      "test tost",
		Title:       "test book",
		Description: "this is book",
	}}

	response, err := client.CreateBook(ctx, request)
	assert.Equal(t, expectedResult.Book.Author, response.Book.Author)
	assert.Equal(t, expectedResult.Book.Title, response.Book.Title)
	assert.Equal(t, expectedResult.Book.Description, response.Book.Description)
	assert.NoError(t, err)
}

func TestBookServiceServer_ReadBook_NotFoundData(t *testing.T) {
	expectedResult := "rpc error: code = NotFound desc = Could not find book with Object Id 5f81e55d9f9ff3494f7a1b10: mongo: no documents in result"

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := bookpb.NewBookServiceClient(conn)
	request := &bookpb.ReadBookReq{Id: "5f81e55d9f9ff3494f7a1b10"}

	_, err = client.ReadBook(ctx, request)
	assert.Equal(t, expectedResult, err.Error())
}

func TestBookServiceServer_UpdateBook_NotFoundData(t *testing.T) {
	expectedResult := "rpc error: code = NotFound desc = Could not find book with supplied ID: mongo: no documents in result"

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := bookpb.NewBookServiceClient(conn)
	request := &bookpb.UpdateBookReq{Book: &bookpb.Book{
		Id:          "5f81e55d9f9ff3494f7a1b10",
		Author:      "test tost",
		Title:       "test book",
		Description: "this is book",
	}}

	_, err = client.UpdateBook(ctx, request)
	assert.Equal(t, expectedResult, err.Error())
}

func TestBookServiceServer_DeleteBook_NotFoundData(t *testing.T) {
	expectedResult := true
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := bookpb.NewBookServiceClient(conn)
	request := &bookpb.DeleteBookReq{Id: "5f81e55d9f9ff3494f7a1b10"}

	response, err := client.DeleteBook(ctx, request)
	assert.Equal(t, expectedResult, response.Success)
	assert.NoError(t, err)
}
