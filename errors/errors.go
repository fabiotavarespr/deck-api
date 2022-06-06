package internalerrors

type ErrNotFound struct {
	Message string
}

func (e ErrNotFound) Error() string {
	return e.Message
}

type ErrInvalidEntry struct {
	Message string
}

func (e ErrInvalidEntry) Error() string {
	return e.Message
}

type ErrInsufficientResources struct {
	Message string
}

func (e ErrInsufficientResources) Error() string {
	return e.Message
}
