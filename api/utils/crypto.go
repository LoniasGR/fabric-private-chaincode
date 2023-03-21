package utils

import (
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

// Creates a mnemonic from randomness
func CreateMnemonic() (string, error) {
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return "", err
	}

	mnemonic, _ := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", err
	}

	return mnemonic, nil
}

// Creates a master key from mnemonic and passphrase
func CreateMasterKey(mnemonic, passphrase string) (string, error) {
	seed := bip39.NewSeed(mnemonic, passphrase)

	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return "", err
	}

	return masterKey.B58Serialize(), nil
}

// Checks if the public key matches the private key
func MasterKeysMatch(masterKeySerialized, publicKeySerialized string) (bool, error) {
	masterKey, err := bip32.B58Deserialize(masterKeySerialized)
	if err != nil {
		return false, err
	}

	publicKey, err := bip32.B58Deserialize(publicKeySerialized)
	if err != nil {
		return false, err
	}

	return masterKey.PublicKey() == publicKey, nil
}

func PubKeyMatchesMnemonic(mnemonic, passphrase, pubkeySerialized string) (bool, error) {
	seed := bip39.NewSeed(mnemonic, passphrase)
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return false, err
	}
	masterKeySerialized := masterKey.PublicKey().B58Serialize()

	return masterKeySerialized == pubkeySerialized, nil
}
