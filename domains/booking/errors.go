package booking

import "github.com/cilloparch/cillop/i18np"

type Errors interface {
	Failed(string) *i18np.Error
	InvalidUUID() *i18np.Error
	InternalError() *i18np.Error
	NotAvailable() *i18np.Error
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
