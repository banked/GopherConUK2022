package main

import (
	"fmt"
	"log"

	"github.com/banked/gopherconuk2022/demo1/github.com/banked/gopherconuk2022/demo1"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	msg := demo1.ErrInsuffcientFunds{
		PaymentId: "123",
		Amount:    int64(1000),
		Currency:  "GBP",
		Provider:  "Bank A",
	}

	b, err := protojson.Marshal(&msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
