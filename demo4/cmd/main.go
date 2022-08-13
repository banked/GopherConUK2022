package main

import (
	"fmt"
	"log"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/banked/gopherconuk2022/demo4"
)

func main() {
	msg := demo4.ErrInsuffcientFunds{
		PaymentId: "123",
		Amount: &demo4.Amount{
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
