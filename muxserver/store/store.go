package store

import (
	"errors"
	"log"
	"webstudy/store/types"
)

type PersistentStorageInterface interface {
	GetUidString() string
}

type storeObj struct{}

var Store PersistentStorageInterface

var uGen types.UidGenerator

// GetUidString generate unique ID as string
func (storeObj) GetUidString() string {
	return uGen.GetStr()
}

func init() {
	Store = storeObj{}
	if err := uGen.Init(uint(1), []byte("1234567890123456")); err != nil {
		log.Println(errors.New("store: failed to init snowflake: " + err.Error()))
	}
}
