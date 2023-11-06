package command

type BookingDetailCmd struct {
	UUID string `params:"uuid" validate:"required,object_id"`
}
