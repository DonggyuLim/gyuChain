package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gyuChain/utils"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

type wallet struct {
	PrivKey *ecdsa.PrivateKey
	PubKey  []byte
}

func NewWallet() *wallet {
	curve := elliptic.P256()
	privKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	utils.PanicErr(err)
	pubKey := append(privKey.X.Bytes(), privKey.Y.Bytes()...)
	return &wallet{privKey, pubKey}
}

type mnemonicWallet struct {
	Mnemonic  string
	Seed      []byte
	MasterKey *bip32.Key
	PublicKey *bip32.Key
}

func NewMnemonic() *mnemonicWallet {
	entropy, err := bip39.NewEntropy(256) //make mnemonic length 24
	utils.PanicErr(err)
	mnemonic, err := bip39.NewMnemonic(entropy)
	fmt.Printf("menonic = %s\n", mnemonic)

	seed := bip39.NewSeed(mnemonic, "Secret passphrase")
	fmt.Printf("seed = %v\n", seed)
	masterKey, err := bip32.NewMasterKey(seed)
	utils.PanicErr(err)

	fmt.Printf("masterKey =%s\n", masterKey)

	publicKey := masterKey.PublicKey()
	fmt.Printf("publicKey = %s\n", publicKey)
	mnemonicWallet := &mnemonicWallet{
		Mnemonic:  mnemonic,
		Seed:      seed,
		MasterKey: masterKey,
		PublicKey: publicKey,
	}

	mb, err := json.MarshalIndent(mnemonicWallet, "", "")
	utils.PanicErr(err)
	err = ioutil.WriteFile("./wallet.json", mb, 0777)
	utils.PanicErr(err)

	return mnemonicWallet
}
