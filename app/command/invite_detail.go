package command

type InviteDetailCmd struct {
	InviteUUID string `params:"uuid" validate:"required,object_id"`
}
