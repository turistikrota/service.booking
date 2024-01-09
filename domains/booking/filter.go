package booking

type FilterEntity struct {
	Locale string `json:"-" query:"-"`
	Query  string `query:"q,omitempty" validate:"omitempty,max=100"`
	State  string `query:"state,omitempty" validate:"omitempty,oneof=canceled not_valid created pay_expired pay_cancelled pay_pending pay_paid pay_refunded"`
}
