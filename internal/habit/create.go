package habit

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type habitCreator interface {
	Add(ctx context.Context, habit Habit) error
}

func Create(ctx context.Context, db habitCreator, h Habit) (Habit, error) {
	h, err := validateAndFillDetails(h)
	if err != nil {
		return Habit{}, err
	}

	err = db.Add(ctx, h)
	if err != nil {
		return Habit{}, fmt.Errorf("cannot save habit: %w", err)
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
