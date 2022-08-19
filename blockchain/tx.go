package blockchain

import "github.com/gyuChain/utils"

type Account struct {
	Nonce   int
	Address string
	Balance int
}

type Tx struct {
	hash      string
	Nonce     int
	From      Account
	to        Account
	Signature string
}

func (t *Tx) getTxId() {
	t.hash = utils.Hash(t)
}

// DB -> Account -> Nonce
func (t *Tx) getNonce() {

}
