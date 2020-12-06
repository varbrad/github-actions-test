package pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		expected := Bitcoin(10)
		actual := wallet.Balance()

		assert.Equal(t, expected, actual)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		expected := Bitcoin(10)
		actual := wallet.Balance()

		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}

		err := wallet.Withdraw(Bitcoin(15))

		assert.Equal(t, Bitcoin(10), wallet.Balance())
		assert.Error(t, err)
	})
}

func TestBitcoin(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		actual := Bitcoin(10).String()
		expected := "10 BTC"

		assert.Equal(t, expected, actual)
	})
}
