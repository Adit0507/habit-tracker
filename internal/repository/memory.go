package repository

import (
	"context"
	"habits/internal/habit"
	"sort"
	"sync"
)

type HabitRepository struct {
	mutex sync.Mutex	
	storage map[habit.ID]habit.Habit
	lgr Logger
}

type Logger interface{
	Logf(format string, args ...any)
}

func New(lgr Logger) *HabitRepository {
	return &HabitRepository{
		lgr: lgr,
		storage: make(map[habit.ID]habit.Habit),
	}
}

func (hr*HabitRepository) Add(_ context.Context, habit habit.Habit) error {
	hr.lgr.Logf("Adding a habit")

	hr.mutex.Lock()
	defer hr.mutex.Unlock()

	hr.storage[habit.ID] = habit

	return nil
}

func (hr*HabitRepository) FindAll(_ context.Context) ([]habit.Habit, error) {
	hr.lgr.Logf("Listing habits...")

	hr.mutex.Lock()
	defer hr.mutex.Unlock()

	habits := make([]habit.Habit, 0)
	for _, h := range hr.storage{
		habits = append(habits, h)
	}

	sort.Slice(habits, func(i, j int) bool {
		return habits[i].CreationTime.Before(habits[j].CreationTime)
	})

	return habits, nil
}