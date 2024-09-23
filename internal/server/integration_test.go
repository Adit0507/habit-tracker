package server

import (
	"context"
	"net"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"habits/api"
	repo "habits/internal/repository"
)

func TestIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	// run server
	grpcServ := newServer(t)
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = grpcServ.Serve(listener)
		require.NoError(t, err)
	}()
	defer func() {
		// terminate the GRPC server
		grpcServ.Stop()
		// when that is done, and no error were caught, we can end this test
		wg.Wait()
	}()

	// create client
	habitsCli, err := newClient(t, listener.Addr().String())
	require.NoError(t, err)

	// add 2 habits
	addHabit(t, habitsCli, nil, "walk in the forest")

	addHabit(t, habitsCli, ptr(3), "read a few pages")

	addHabitWithError(t, habitsCli, 5, "  	  ", codes.InvalidArgument)

	// check that the 2 habits are present
	listHabitsMatches(t, habitsCli, []*api.Habit{
		{
			Name:            "walk in the forest",
			WeeklyFrequency: 1,
		},
		{
			Name:            "read a few pages",
			WeeklyFrequency: 3,
		},
	})
}

func newServer(t *testing.T) *grpc.Server {
	t.Helper()
	// Our t variable implements the Logger interfaces, as it exposes Logf(...).
	s := New(repo.New(t), t)

	return s.registerGRPCServer()
}

func newClient(t *testing.T, serverAddress string) (api.HabitsClient, error) {
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(serverAddress, creds)
	require.NoError(t, err)

	return api.NewHabitsClient(conn), nil
}

func addHabit(t *testing.T, habitsCli api.HabitsClient, freq *int32, name string) {
	_, err := habitsCli.CreateHabit(context.Background(), &api.CreateHabitRequest{
		Name:            name,
		WeeklyFrequency: freq,
	})
	require.NoError(t, err)
}

func ptr(i int32) *int32 {
	return &i
}

func addHabitWithError(t *testing.T, habitsCli api.HabitsClient, freq int32, name string, statusCode codes.Code) {
	_, err := habitsCli.CreateHabit(context.Background(), &api.CreateHabitRequest{
		Name:            name,
		WeeklyFrequency: &freq,
	})
	statusErr, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, statusCode, statusErr.Code())
}

func listHabitsMatches(t *testing.T, habitsCli api.HabitsClient, expected []*api.Habit) {
	list, err := habitsCli.ListHabits(context.Background(), &api.ListHabitsRequest{})
	require.NoError(t, err)

	for i := range list.Habits {
		assert.NotEqual(t, "", list.Habits[i].Id)
		list.Habits[i].Id = "" // generated
	}
	assert.Equal(t, list.Habits, expected)
}