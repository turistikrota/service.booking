package command

type PostDetailCmd struct {
	PostUUID string `params:"postUUID" validate:"required,object_id"`
}
