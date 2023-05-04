package helper

import (
	"net/http"
)

const (
	ErrPlatformNotSupported  = "your platform is unsupported! i can't clear terminal screen :("
	ErrDbUrlNotExist         = "database URL not found"
	ErrEnvNotFound           = ".env file not found"
	ErrContactNameNotValid   = "name yang dimasukkan tidak valid"
	ErrContactNoTelpNotValid = "no_telp yang dimasukkan tidak valid"
	ErrContactIdNotValid     = "contact id yang dimasukkan tidak valid"

	ErrContactNotFound = "contact not found"
)

func HandleAppError(err error) (int, string) {
	switch e := err.(type) {
	case *AppError:
		switch e.Message {
		case ErrContactNotFound:
			return http.StatusNotFound, err.Error()
		default:
			return http.StatusInternalServerError, err.Error()
		}

	default:
		return http.StatusInternalServerError, err.Error()
	}
}

type AppError struct {
	Message string
}

func NewAppError(message string) *AppError {
	return &AppError{
		Message: message,
	}
}

func (e *AppError) Error() string {
	return e.Message
}
