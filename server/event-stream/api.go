package event_stream

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/turistikrota/service.booking/app/command"
)

func (s srv) OnBookingValidationSucceed(data []byte) {
	fmt.Println("OnBookingValidationSucceed")
	cmd := command.BookingValidationSucceedCmd{}
	err := json.Unmarshal(data, &cmd)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", cmd)
	_, _ = s.app.Commands.BookingValidationSucceed(context.Background(), cmd)
}

func (s srv) OnBookingValidationFail(data []byte) {
	fmt.Println("OnBookingValidationFail")
	cmd := command.BookingValidationFailedCmd{}
	err := json.Unmarshal(data, &cmd)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", cmd)
	_, _ = s.app.Commands.BookingValidationFailed(context.Background(), cmd)
}
