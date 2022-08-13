package main

import (
	"fmt"
	"log"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"

	paymentv1 "github.com/banked/gopherconuk2022/demo5/banked/payment/v1"
	bankedv1 "github.com/banked/gopherconuk2022/demo5/banked/v1"
)

func main() {
	msg := paymentv1.ErrInsuffcientFunds{
		PaymentId: "123",
		Amount: &bankedv1.Amount{
			Cents:    int64(10000),
			Currency: "GBP",
		},
		Provider:    "Bank A",
		AttemptedAt: timestamppb.New(time.Now().UTC()),
	}

	b, err := protojson.Marshal(&msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
