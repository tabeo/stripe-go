package connectiontoken

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/tabeo/stripe-go"
	_ "github.com/tabeo/stripe-go/testing"
)

func TestTerminalConnectionTokenNew(t *testing.T) {
	connectiontoken, err := New(&stripe.TerminalConnectionTokenParams{})
	assert.Nil(t, err)
	assert.NotNil(t, connectiontoken)
}
