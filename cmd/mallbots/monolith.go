package main

import (
	"database/sql"
	"eda-in-go/internal/config"
	"eda-in-go/internal/monolith"
	"eda-in-go/waiter"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type app struct {
	cfg     config.AppConfig
	db      *sql.DB
	logger  zerolog.Logger
	modules []monolith.Module
	mux     *chi.Mux
	rpc     *grpc.Server
	waiter  waiter.Waiter
}

func (a *app) Config() config.AppConfig {
	return a.cfg
}

func (a *app) DB() *sql.DB {
	return a.db
}

func (a *app) Logger() zerolog.Logger {
	return a.logger
}

func (a *app) Mux() *chi.Mux {
	return a.mux
}

func (a *app) RPC() *grpc.Server {
	return a.rpc
}

func (a *app) Waiter() waiter.Waiter {
	return a.waiter
}
