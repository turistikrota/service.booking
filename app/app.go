package app

import "github.com/turistikrota/service.booking/app/command"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	BookingCreate            command.BookingCreateHandler
	BookingValidationSucceed command.BookingValidationSucceedHandler
	BookingValidationFailed  command.BookingValidationFailedHandler
}

type Queries struct {
}
