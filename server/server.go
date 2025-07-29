package server

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/leninner/shared/config"
)

func Serve(app *config.Application, handler http.Handler) error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Config.Port),
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(app.Logger.Handler(), slog.LevelError),
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		app.Logger.Info("shutting down server", "signal", s.String())

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		app.WG.Wait()
		shutdownError <- srv.Shutdown(ctx)
	}()

	app.Logger.Info("starting server", "addr", srv.Addr, "env", app.Config.Env)

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	app.Logger.Info("stopped server", "addr", srv.Addr)

	return nil
} 

func OpenDB(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DB.DSN)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.DB.MaxOpenConns)
	db.SetMaxIdleConns(cfg.DB.MaxIdleConns)
	db.SetConnMaxIdleTime(cfg.DB.MaxIdleTime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}

		log.Println("database connection pool established")

	return db, nil
}

