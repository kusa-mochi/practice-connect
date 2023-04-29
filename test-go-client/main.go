package main

import (
	"context"
	"log"
	"net/http"
	v1 "proto_example/gen/my_interface/v1"
	v1connect "proto_example/gen/my_interface/v1/my_interfacev1connect"

	"github.com/bufbuild/connect-go"
)

func TestTestRpc1() {
	log.Println("TestRpc1 start")
	defer log.Println("TestRpc1 end")
	client := v1connect.NewTestService1Client(
		http.DefaultClient,
		"http://localhost:8080",
	)
	res, err := client.TestRpc1(
		context.Background(),
		connect.NewRequest(&v1.TestRequest1{RequestContent: "tesutesu1"}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res.Msg.ResponseContent)
}

func TestTestRpc1_1() {
	log.Println("TestRpc1_1 start")
	defer log.Println("TestRpc1_1 end")
	client := v1connect.NewTestService1Client(
		http.DefaultClient,
		"http://localhost:8080",
	)
	res, err := client.TestRpc1_1(
		context.Background(),
		connect.NewRequest(&v1.TestRequest2{RequestValue: 100}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res.Msg.ResponseValue)
}

func TestTestRpc1_2() {
	log.Println("TestRpc1_2 start")
	defer log.Println("TestRpc1_2 end")
	client := v1connect.NewTestService1Client(
		http.DefaultClient,
		"http://localhost:8080",
	)
	res, err := client.TestRpc1_2(
		context.Background(),
		connect.NewRequest(&v1.CommonMessage{Content: "tesu", EnumValue: v1.TestEnum_Value1}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res.Msg.ResponseContent)
}

func TestTestRpc2() {
	log.Println("TestRpc2 start")
	defer log.Println("TestRpc2 end")
	client := v1connect.NewTestService2Client(
		http.DefaultClient,
		"http://localhost:8080",
	)
	res, err := client.TestRpc2(
		context.Background(),
		connect.NewRequest(&v1.TestRequest2{RequestValue: 200}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res.Msg.ResponseValue)
}

func TestTestRpc2_1() {
	log.Println("TestRpc2_1 start")
	defer log.Println("TestRpc2_1 end")
	client := v1connect.NewTestService2Client(
		http.DefaultClient,
		"http://localhost:8080",
	)
	res, err := client.TestRpc2_1(
		context.Background(),
		connect.NewRequest(&v1.CommonMessage{Content: "tesu2_1!", EnumValue: v1.TestEnum_Value2}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res.Msg.ResponseValue)
}

func main() {
	TestTestRpc1()
	TestTestRpc1_1()
	TestTestRpc1_2()
	TestTestRpc2()
	TestTestRpc2_1()
}
