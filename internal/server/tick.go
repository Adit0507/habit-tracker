package server

import (
	"context"
	"errors"
	"fmt"
	"habits/api"
	"habits/internal/habit"
	r "habits/internal/repository"

	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validateTickHabitRequest(req *api.TickHabitRequest) error {
	switch {
	case req == nil:
		return fmt.Errorf("empty request")
	case req.HabitId == "":
		return fmt.Errorf("missing habit ID")
	}
	return nil
}

func (s *Server) TickHabit(ctx context.Context, req *api.TickHabitRequest) (*api.TickHabitResponse, error) {
	s.lgr.Logf("Tick request received: %s", req)

	err := validateTickHabitRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request: "+err.Error())
	}

	var t time.Time
	if req.Timestamp == nil {
		t = time.Now()
	}

	err = habit.Tick(ctx, s.db, s.db, habit.ID(req.HabitId), t)
	if err != nil {
		switch {
		case errors.Is(err, r.ErrNotFound):
			return nil, status.Errorf(codes.NotFound, "couldn't find habit %q in repository", req.HabitId)
		default:
			return nil, status.Errorf(codes.Internal, "cannot tick habit %q: %s", req.HabitId, err.Error())
		}
	}


	return &api.TickHabitResponse{}, nil
}
