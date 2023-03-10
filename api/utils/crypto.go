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
func CreateMasterKey(mnemonic, passphrase string) ([]byte, error) {
	seed := bip39.NewSeed(mnemonic, passphrase)

	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, err
	}

	keySerialized, err := masterKey.Serialize()
	if err != nil {
		return nil, err
	}
	return keySerialized, nil
}

// Checks if the public key matches the private key
func MasterKeysMatch(masterKeySerialized, publicKeySerialized []byte) (bool, error) {
	masterKey, err := bip32.Deserialize(masterKeySerialized)
	if err != nil {
		return false, err
	}

	publicKey, err := bip32.Deserialize(publicKeySerialized)
	if err != nil {
		return false, err
	}

	return masterKey.PublicKey() == publicKey, nil
}

func PubKeyMatchesMnemonic(mnemonic, passphrase string, pubkeySerialized []byte) (bool, error) {
	seed := bip39.NewSeed(mnemonic, passphrase)
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return false, err
	}

	publicKey, err := bip32.Deserialize(pubkeySerialized)
	if err != nil {
		return false, err
	}

	return masterKey.PublicKey() == publicKey, nil
}
