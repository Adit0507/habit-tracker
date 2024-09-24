package repository

type Error string

func (e Error) Error() string {
	return string(e)
}

const ErrNotFound = Error("habit not found")