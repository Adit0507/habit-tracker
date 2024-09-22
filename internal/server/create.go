package server

import (
	"context"
	"habits/api"
)

func (s*Server) CreateHabit(_ context.Context, request *api.CreateHabitRequest) (*api.CreateHabitResponse, error) {
	s.lgr.Logf("Create Habit Request received: %s", request)

	return &api.CreateHabitResponse{
		Habit: &api.Habit{
			Name: "aa",
			WeeklyFrequency: 1,
		},
	}, nil
}