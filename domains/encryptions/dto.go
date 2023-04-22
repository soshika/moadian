package encrytpions

import (
	"crypto/rand"
)

type AESGCM struct {
	KeySize int64 `json:"key_size"`
}

func (eg *AESGCM) getAESKey() string {

	key := make([]byte, eg.KeySize)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}

	return string(key)
}
