package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dexguitar/gotododbservice/internal/app"
)

func main() {
	application := app.New("postgres://postgres:qwerty@localhost:5432/gotodo?sslmode=disable")

	go application.GRPCSrv.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	sign := <-stop

	fmt.Println("stopping sso server", sign.String())

	application.GRPCSrv.Stop()
}
