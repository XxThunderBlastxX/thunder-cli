package service

import "github.com/99designs/keyring"

func NewKeyRingService() (keyring.Keyring, error) {
	k, err := keyring.Open(keyring.Config{
		AllowedBackends: []keyring.BackendType{keyring.KeychainBackend},
		ServiceName:     "thunder",
	})
	if err != nil {
		return nil, err
	}

	return k, nil
}
