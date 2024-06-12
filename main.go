package main

import (
	_ "aiplus/docs"
	"aiplus/internal/app/controller"
	"aiplus/internal/app/repository"
	"aiplus/internal/app/service"
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/subosito/gotenv"
	"golang.org/x/sync/errgroup"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"syscall"
)

//go:generate swag init --parseInternal --pd
func main() {
	_ = gotenv.OverLoad()

	dbConn, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")))
	if err != nil {
		log.Fatal(err)
	}

	r, err := repository.NewRepository(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	s := service.NewService(r)

	ctrl := controller.NewController(s).Build()

	group, ctx := errgroup.WithContext(context.Background())
	{
		group.Go(func() error {
			err = ctrl.Start(":8080")
			if err != nil {
				return err
			}
			return ctrl.Shutdown(ctx)
		})
	}
	{
		group.Go(
			func() error {
				cs := make(chan os.Signal, 1)
				signal.Notify(cs, syscall.SIGINT, syscall.SIGTERM)
				select {
				case sig := <-cs:
					return fmt.Errorf("received signal %s", sig)
				}
			},
		)
	}

	log.Warn("exit", group.Wait())
}
