package booking

import "github.com/cilloparch/cillop/i18np"

type ValidationError struct {
	Field   string  `json:"field"`
	Message string  `json:"message"`
	Params  i18np.P `json:"params"`
}

type Errors interface {
	Failed(string) *i18np.Error
	InvalidUUID() *i18np.Error
	InternalError() *i18np.Error
	NotAvailable() *i18np.Error
	OnlyAdminCanDoThisAction() *i18np.Error
	StartDateAfterEndDate() *i18np.Error
	StartDateBeforeNow() *i18np.Error
	NotCancelable() *i18np.Error
	NotFound() *i18np.Error
}

type errors struct{}

func NewErrors() Errors {
	return &errors{}
}

func (e *errors) Failed(action string) *i18np.Error {
	return i18np.NewError(I18nMessages.Failed, i18np.P{
		"Action": action,
	})
}

func (e *errors) InvalidUUID() *i18np.Error {
	return i18np.NewError(I18nMessages.InvalidUUID)
}

func (e *errors) InternalError() *i18np.Error {
	return i18np.NewError(I18nMessages.InternalError)
}

func (e *errors) NotAvailable() *i18np.Error {
	return i18np.NewError(I18nMessages.NotAvailable)
}

func (e *errors) StartDateAfterEndDate() *i18np.Error {
	return i18np.NewError(I18nMessages.StartDateAfterEndDate)
}

func (e *errors) StartDateBeforeNow() *i18np.Error {
	return i18np.NewError(I18nMessages.StartDateBeforeNow)
}

func (e *errors) OnlyAdminCanDoThisAction() *i18np.Error {
	return i18np.NewError(I18nMessages.OnlyAdminCanDoThisAction)
}

func (e *errors) NotCancelable() *i18np.Error {
	return i18np.NewError(I18nMessages.NotCancelable)
}

func (e *errors) NotFound() *i18np.Error {
	return i18np.NewError(I18nMessages.NotFound)
}