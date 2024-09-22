package habit

import "time"

type ID string

type Name string

type WeeklyFrequency uint

type Habit struct {
	ID              ID
	Name            Name
	WeeklyFrequency WeeklyFrequency
	CreationTime    time.Time
}