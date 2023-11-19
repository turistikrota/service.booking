package command

type ListingDetailCmd struct {
	ListingUUID string `params:"listingUUID" validate:"required,object_id"`
}
