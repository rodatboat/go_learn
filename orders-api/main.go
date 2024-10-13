package main

import (
	"context"
	"fmt"
	"go_learn/orders-api/application"
	"os"
	"os/signal"
)

func main() {
	app := application.New(application.LoadConfig())

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel() // defer means to call it at the end of the current function (last in, first out)

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
