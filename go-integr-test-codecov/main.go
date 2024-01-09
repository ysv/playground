package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	e := echo.New()
	e.HideBanner = true

	e.GET("/time", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct {
			Time string `json:"time"`
		}{
			time.Now().UTC().Format(time.RFC3339),
		})
	})

	e.GET("/version", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct {
			Version string `json:"version"`
		}{
			"0.0.1",
		})
	})

	errGrp, ctx := errgroup.WithContext(ctx)
	errGrp.Go(func() error {
		port := ":8080"
		err := e.Start(port)

		if err != nil && err != http.ErrServerClosed {
			return err
		}

		return nil
	})

	errGrp.Go(func() error {
		<-ctx.Done()
		return e.Shutdown(context.Background())
	})

	lo.Must0(errGrp.Wait())
}
