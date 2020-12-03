package stripe

import (
	"encoding/json"

	"github.com/tabeo/stripe-go/form"
)

// SubscriptionScheduleEndBehavior describe what happens to a schedule when it ends.
type SubscriptionScheduleEndBehavior string

// List of values that SubscriptionScheduleEndBehavior can take.
const (
	SubscriptionScheduleEndBehaviorCancel  SubscriptionScheduleEndBehavior = "cancel"
	SubscriptionScheduleEndBehaviorRelease SubscriptionScheduleEndBehavior = "release"
)

// SubscriptionScheduleStatus is the list of allowed values for the schedule's status.
type SubscriptionScheduleStatus string

// List of values that SubscriptionScheduleStatus can take.
const (
	SubscriptionScheduleStatusActive    SubscriptionScheduleStatus = "active"
	SubscriptionScheduleStatusCanceled  SubscriptionScheduleStatus = "canceled"
	SubscriptionScheduleStatusCompleted SubscriptionScheduleStatus = "completed"
	SubscriptionScheduleStatusPastDue   SubscriptionScheduleStatus = "not_started"
	SubscriptionScheduleStatusTrialing  SubscriptionScheduleStatus = "released"
)

// SubscriptionScheduleInvoiceSettingsParams is a structure representing the parameters allowed to
// control invoice settings on invoices associated with a subscription schedule.
type SubscriptionScheduleInvoiceSettingsParams struct {
	DaysUntilDue *int64 `form:"days_until_due"`
}

// SubscriptionScheduleDefaultSettingsParams is the set of parameters
// representing the subscription schedule’s default settings.
type SubscriptionScheduleDefaultSettingsParams struct {
	Params               `form:"*"`
	BillingThresholds    *SubscriptionBillingThresholdsParams       `form:"billing_thresholds"`
	CollectionMethod     *string                                    `form:"collection_method"`
	DefaultPaymentMethod *string                                    `form:"default_payment_method"`
	InvoiceSettings      *SubscriptionScheduleInvoiceSettingsParams `form:"invoice_settings"`
}

// SubscriptionSchedulePhaseAddInvoiceItemPriceDataRecurringParams is a structure representing the
// parameters to create an inline recurring price for a given invoice item.
type SubscriptionSchedulePhaseAddInvoiceItemPriceDataRecurringParams struct {
	AggregateUsage  *string `form:"aggregate_usage"`
	Interval        *string `form:"interval"`
	IntervalCount   *int64  `form:"interval_count"`
	TrialPeriodDays *int64  `form:"trial_period_days"`
	UsageType       *string `form:"usage_type"`
}

// SubscriptionSchedulePhaseAddInvoiceItemPriceDataParams is a structure representing the parameters
// to create an inline price for a given invoice item.
type SubscriptionSchedulePhaseAddInvoiceItemPriceDataParams struct {
	Currency          *string                                                          `form:"currency"`
	Product           *string                                                          `form:"product"`
	Recurring         *SubscriptionSchedulePhaseAddInvoiceItemPriceDataRecurringParams `form:"recurring"`
	UnitAmount        *int64                                                           `form:"unit_amount"`
	UnitAmountDecimal *float64                                                         `form:"unit_amount_decimal,high_precision"`
}

// SubscriptionSchedulePhaseAddInvoiceItemParams is a structure representing the parameters allowed to control
// the invoice items to add at the start of a phase.
type SubscriptionSchedulePhaseAddInvoiceItemParams struct {
	Price     *string                     `form:"price"`
	PriceData *InvoiceItemPriceDataParams `form:"price_data"`
	Quantity  *int64                      `form:"quantity"`
}

// SubscriptionSchedulePhaseItemParams is a structure representing the parameters allowed to control
// a specic plan on a phase on a subscription schedule.
type SubscriptionSchedulePhaseItemParams struct {
	BillingThresholds *SubscriptionItemBillingThresholdsParams `form:"billing_thresholds"`
	Plan              *string                                  `form:"plan"`
	Price             *string                                  `form:"price"`
	PriceData         *SubscriptionItemPriceDataParams         `form:"price_data"`
	Quantity          *int64                                   `form:"quantity"`
	TaxRates          []*string                                `form:"tax_rates"`
}

// SubscriptionSchedulePhaseParams is a structure representing the parameters allowed to control
// a phase on a subscription schedule.
type SubscriptionSchedulePhaseParams struct {
	AddInvoiceItems       []*SubscriptionSchedulePhaseAddInvoiceItemParams `form:"add_invoice_items"`
	ApplicationFeePercent *int64                                           `form:"application_fee_percent"`
	BillingThresholds     *SubscriptionBillingThresholdsParams             `form:"billing_thresholds"`
	CollectionMethod      *string                                          `form:"collection_method"`
	Coupon                *string                                          `form:"coupon"`
	DefaultPaymentMethod  *string                                          `form:"default_payment_method"`
	DefaultTaxRates       []*string                                        `form:"default_tax_rates"`
	EndDate               *int64                                           `form:"end_date"`
	InvoiceSettings       *SubscriptionScheduleInvoiceSettingsParams       `form:"invoice_settings"`
	Iterations            *int64                                           `form:"iterations"`
	Plans                 []*SubscriptionSchedulePhaseItemParams           `form:"plans"`
	ProrationBehavior     *string                                          `form:"proration_behavior"`
	StartDate             *int64                                           `form:"start_date"`
	Trial                 *bool                                            `form:"trial"`
	TrialEnd              *int64                                           `form:"trial_end"`

	// This parameter is deprecated and we recommend that you use TaxRates instead.
	TaxPercent *float64 `form:"tax_percent"`
}

// SubscriptionScheduleParams is the set of parameters that can be used when creating or updating a
// subscription schedule.
type SubscriptionScheduleParams struct {
	Params            `form:"*"`
	Customer          *string                                    `form:"customer"`
	DefaultSettings   *SubscriptionScheduleDefaultSettingsParams `form:"default_settings"`
	EndBehavior       *string                                    `form:"end_behavior"`
	FromSubscription  *string                                    `form:"from_subscription"`
	ProrationBehavior *string                                    `form:"proration_behavior"`
	Phases            []*SubscriptionSchedulePhaseParams         `form:"phases"`
	StartDate         *int64                                     `form:"start_date"`
	StartDateNow      *bool                                      `form:"-"` // See custom AppendTo

	// TODO remove in the next major version
	// This propery is considered deprecated. Use ProrationBehavior instead
	Prorate *bool `form:"prorate"`
}

// AppendTo implements custom encoding logic for SubscriptionScheduleParams so that the special
// "now" value for start_date can be implemented (they're otherwise timestamps rather than strings).
func (p *SubscriptionScheduleParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.StartDateNow) {
		body.Add(form.FormatKey(append(keyParts, "start_date")), "now")
	}
}

// SubscriptionScheduleCancelParams is the set of parameters that can be used when canceling a
// subscription schedule.
type SubscriptionScheduleCancelParams struct {
	Params     `form:"*"`
	InvoiceNow *bool `form:"invoice_now"`
	Prorate    *bool `form:"prorate"`
}

// SubscriptionScheduleReleaseParams is the set of parameters that can be used when releasing a
// subscription schedule.
type SubscriptionScheduleReleaseParams struct {
	Params             `form:"*"`
	PreserveCancelDate *bool `form:"preserve_cancel_date"`
}

// SubscriptionScheduleListParams is the set of parameters that can be used when listing
// subscription schedules.
type SubscriptionScheduleListParams struct {
	ListParams       `form:"*"`
	CanceledAt       int64             `form:"canceled_at"`
	CanceledAtRange  *RangeQueryParams `form:"canceled_at"`
	CompletedAt      int64             `form:"completed_at"`
	CompletedAtRange *RangeQueryParams `form:"completed_at"`
	Created          int64             `form:"created"`
	CreatedRange     *RangeQueryParams `form:"created"`
	Customer         string            `form:"customer"`
	ReleasedAt       int64             `form:"released_at"`
	ReleasedAtRange  *RangeQueryParams `form:"released_at"`
	Scheduled        *bool             `form:"scheduled"`
}

// SubscriptionScheduleCurrentPhase is a structure the current phase's period.
type SubscriptionScheduleCurrentPhase struct {
	EndDate   int64 `json:"end_date"`
	StartDate int64 `json:"start_date"`
}

// SubscriptionScheduleInvoiceSettings is a structure representing the settings applied to invoices
// associated with a subscription schedule.
type SubscriptionScheduleInvoiceSettings struct {
	DaysUntilDue int64 `json:"days_until_due"`
}

// SubscriptionScheduleDefaultSettings is a structure representing the
// subscription schedule’s default settings.
type SubscriptionScheduleDefaultSettings struct {
	BillingThresholds    *SubscriptionBillingThresholds       `json:"billing_thresholds"`
	CollectionMethod     SubscriptionCollectionMethod         `json:"collection_method"`
	DefaultPaymentMethod *PaymentMethod                       `json:"default_payment_method"`
	InvoiceSettings      *SubscriptionScheduleInvoiceSettings `json:"invoice_settings"`
}

// SubscriptionSchedulePhaseAddInvoiceItem represents the invoice items to add when the phase starts.
type SubscriptionSchedulePhaseAddInvoiceItem struct {
	Price    *Price `json:"price"`
	Quantity int64  `json:"quantity"`
}

// SubscriptionSchedulePhaseItem represents plan details for a given phase
type SubscriptionSchedulePhaseItem struct {
	BillingThresholds *SubscriptionItemBillingThresholds `json:"billing_thresholds"`
	Plan              *Plan                              `json:"plan"`
	Price             *Price                             `json:"price"`
	Quantity          int64                              `json:"quantity"`
	TaxRates          []*TaxRate                         `json:"tax_rates"`
}

// SubscriptionSchedulePhase is a structure a phase of a subscription schedule.
type SubscriptionSchedulePhase struct {
	AddInvoiceItems       []*SubscriptionSchedulePhaseAddInvoiceItem `json:"add_invoice_items"`
	ApplicationFeePercent float64                                    `json:"application_fee_percent"`
	BillingThresholds     *SubscriptionBillingThresholds             `json:"billing_thresholds"`
	CollectionMethod      SubscriptionCollectionMethod               `json:"collection_method"`
	Coupon                *Coupon                                    `json:"coupon"`
	DefaultPaymentMethod  *PaymentMethod                             `json:"default_payment_method"`
	DefaultTaxRates       []*TaxRate                                 `json:"default_tax_rates"`
	EndDate               int64                                      `json:"end_date"`
	InvoiceSettings       *SubscriptionScheduleInvoiceSettings       `json:"invoice_settings"`
	Plans                 []*SubscriptionSchedulePhaseItem           `json:"plans"`
	StartDate             int64                                      `json:"start_date"`
	TrialEnd              int64                                      `json:"trial_end"`

	// This field is deprecated and we recommend that you use TaxRates instead.
	TaxPercent float64 `json:"tax_percent"`
}

// SubscriptionSchedule is the resource representing a Stripe subscription schedule.
type SubscriptionSchedule struct {
	APIResource
	CanceledAt           int64                                `json:"canceled_at"`
	CompletedAt          int64                                `json:"completed_at"`
	Created              int64                                `json:"created"`
	CurrentPhase         *SubscriptionScheduleCurrentPhase    `json:"current_phase"`
	Customer             *Customer                            `json:"customer"`
	DefaultSettings      *SubscriptionScheduleDefaultSettings `json:"default_settings"`
	EndBehavior          SubscriptionScheduleEndBehavior      `json:"end_behavior"`
	ID                   string                               `json:"id"`
	Livemode             bool                                 `json:"livemode"`
	Metadata             map[string]string                    `json:"metadata"`
	Object               string                               `json:"object"`
	Phases               []*SubscriptionSchedulePhase         `json:"phases"`
	ReleasedSubscription *Subscription                        `json:"released_subscription"`
	Status               SubscriptionScheduleStatus           `json:"status"`
	Subscription         *Subscription                        `json:"subscription"`
}

// SubscriptionScheduleList is a list object for subscription schedules.
type SubscriptionScheduleList struct {
	APIResource
	ListMeta
	Data []*SubscriptionSchedule `json:"data"`
}

// UnmarshalJSON handles deserialization of a SubscriptionSchedule.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (s *SubscriptionSchedule) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		s.ID = id
		return nil
	}

	type schedule SubscriptionSchedule
	var v schedule
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*s = SubscriptionSchedule(v)
	return nil
}
