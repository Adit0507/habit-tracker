package server

import (
	"context"
	"fmt"
	"habits/api"
	"habits/internal/habit"
	"net"
	_ "net/http/pprof"
	"strconv"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	db  Repository
}

func New(repo Repository, lgr Logger) *Server {
	return &Server{
		db:  repo,
		lgr: lgr,
	}
}

func (s *Server) registerGRPCServer() *grpc.Server {
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(timerInterceptor(s.lgr)))
	api.RegisterHabitsServer(grpcServer, s)
	reflection.Register(grpcServer) 
	return grpcServer
}

func (s *Server) ListenAndServe(ctx context.Context ,port int) error {
	const addr = "127.0.0.1"
	// combines host & port into a network address
	listener, err := net.Listen("tcp", net.JoinHostPort(addr, strconv.Itoa(port)))
	if err != nil {
		return fmt.Errorf("unable to listen to tcp port %d: %w", port, err)
	}

	grpcServer := s.registerGRPCServer()
	s.lgr.Logf("gRPC server started and listening to port %d", port)

	// Use a channel to report errors from the gRPC server back to
	errChan := make(chan error)
	g := errgroup.Group{}
	defer func() {
		err := g.Wait()
		if err != nil {
			errChan <- fmt.Errorf("error while serving: %w", err)
		}
		close(errChan)
	}()

	// ListenAndServe to the port. This will only return when something kills or stops the server.
	g.Go(func() error {
		// This goroutine will be killed when the context is ended at the end of this function.
		err := grpcServer.Serve(listener)
		if err != nil {
			s.lgr.Logf("error while serving gRPC: %s", err)

			return fmt.Errorf("gRPC server error: %w", err)
		}

		return nil
	})

	select {
	case <-ctx.Done():
		// Stop or GracefulStop was called, no reason to be alarmed.
		s.lgr.Logf("Shutting down grpc server: %s", ctx.Err())
	case err = <-errChan:
		s.lgr.Logf("unable to serve: %w", err)
	}

	grpcServer.GracefulStop()
	_ = listener.Close()
	return nil
}
