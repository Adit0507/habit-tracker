package server

import (
	"context"
	"fmt"
	"habits/api"
	"habits/internal/habit"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

type Logger interface {
	Logf(format string, args ...any)
}

type Repository interface {
	Add(ctx context.Context, habit habit.Habit) error
	FindAll(ctx context.Context) ([]habit.Habit, error)
}

type Server struct {
	api.UnimplementedHabitsServer
	lgr Logger
	db Repository
}

func New(repo Repository ,lgr Logger) *Server {
	return &Server{
		db: repo,
		lgr: lgr,
	}
}

func (s *Server) ListenAndServe(port int) error {
	const addr = "127.0.0.1"
	// combines host & port into a network address
	listener, err := net.Listen("tcp", net.JoinHostPort(addr, strconv.Itoa(port)))
	if err != nil {
		return fmt.Errorf("unable to listen to tcp port %d: %w", port, err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterHabitsServer(grpcServer, s)

	s.lgr.Logf("starting server on port: %d\n", port)

	err = grpcServer.Serve(listener)
	if err != nil {
		return fmt.Errorf("error while listening: %w", err)
	}

	return nil
}
