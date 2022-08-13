package main

import (
	"fmt"
	"log"
	"time"

	"github.com/banked/gopherconuk2022/demo2/github.com/banked/gopherconuk2022/demo2"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	msg := demo2.ErrInsuffcientFunds{
		PaymentId: "123",
		Amount: &demo2.Amount{
			Currency: "GBP",
			Cents:    1000,
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
