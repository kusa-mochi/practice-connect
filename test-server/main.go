package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	v1 "proto_example/gen/my_interface/v1"                              // generated by protoc-gen-go
	v1connect "proto_example/gen/my_interface/v1/my_interfacev1connect" // generated by protoc-gen-connect-go
)

type TestService1 struct{}

func (s *TestService1) TestRpc1(
	ctx context.Context,
	req *connect.Request[v1.TestRequest1],
) (*connect.Response[v1.TestResponse1], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&v1.TestResponse1{
		ResponseContent: fmt.Sprintf("Hello, %s!", req.Msg.RequestContent),
	})
	return res, nil
}

func (s *TestService1) TestRpc1_1(
	ctx context.Context,
	req *connect.Request[v1.TestRequest2],
) (*connect.Response[v1.TestResponse2], error) {
	res := connect.NewResponse(&v1.TestResponse2{
		ResponseValue: req.Msg.RequestValue + 1,
	})
	return res, nil
}

func (s *TestService1) TestRpc1_2(
	ctx context.Context,
	req *connect.Request[v1.CommonMessage],
) (*connect.Response[v1.TestResponse1], error) {
	log.Printf("enum value:%v\n", req.Msg.EnumValue)
	res := connect.NewResponse(&v1.TestResponse1{
		ResponseContent: "TestRpc1_2 response",
	})
	return res, nil
}

type TestService2 struct{}

func main() {
	service1 := &TestService1{}
	mux := http.NewServeMux()
	path, handler := v1connect.NewTestService1Handler(service1)
	mux.Handle(path, handler)
	log.Println("listening...")
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
