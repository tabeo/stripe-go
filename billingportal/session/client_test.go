package session

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/tabeo/stripe-go"
	_ "github.com/tabeo/stripe-go/testing"
)

func TestBillingPortalSessionNew(t *testing.T) {
	session, err := New(&stripe.BillingPortalSessionParams{
		Customer:  stripe.String("cus_123"),
		ReturnURL: stripe.String("https://stripe.com/return"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, session)
}
