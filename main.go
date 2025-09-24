package main

import (
	"context"
	"fmt"

	"github.com/andruixxd31/orders-api/application"
)

func main() {
	app := application.New()

	ctx := context.Background()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("error ", err)
	}
}
