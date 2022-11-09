package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/eferroni/gointensivo/internal/order/infra/database"
	"github.com/eferroni/gointensivo/internal/order/usecase"
	"github.com/eferroni/gointensivo/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"

	//sqlite3
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := database.NewOrderRepository(db)
	uc := usecase.CalculateFinalPriceUseCase{OrderRepository: repository}

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	out := make(chan amqp.Delivery) // channel
	go rabbitmq.Consume(ch, out) // T2

	for msg := range out {
		var inputDto usecase.OrderInputDto
		err := json.Unmarshal(msg.Body, &inputDto)
		if err != nil {
			panic(err)
		}
		outputDto, err := uc.Execute(inputDto)
		if err != nil {
			panic(err)
		}
		msg.Ack(false)
		fmt.Println(outputDto)
		time.Sleep(500 * time.Millisecond)
	}
}