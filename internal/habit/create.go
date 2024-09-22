package habit

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
)

func Create(ctx context.Context, h Habit) (Habit, error) {
	h, err := validateAndFillDetails(h)
	if err != nil {
		return Habit{}, err
	}

	return h, nil
}

func validateAndFillDetails(h Habit) (Habit, error) {
	h.Name = Name(strings.TrimSpace(string(h.Name)))
	if h.Name == "" {
		return Habit{}, InvalidInputError{field: "name", reason: "cannot be empty"}
	}

	if h.WeeklyFrequency == 0 {
		h.WeeklyFrequency = 1
	}

	if h.ID == "" {
		h.ID = ID(uuid.NewString())
	}

	if h.CreationTime.Equal(time.Time{}) {
		h.CreationTime = time.Now()
	}

	return h, nil
}
