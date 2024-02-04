package bin

import (
	"errors"
	"fmt"
	"hash"
	"reflect"

	"github.com/minio/highwayhash"
)

type Types struct {
	types  map[reflect.Type]uint64
	coders map[uint64]Type
	hash   hash.Hash64
}

func NewTypes(key []byte) (t *Types, err error) {
	hash, err := highwayhash.New64(key)

	if err != nil {
		return
	}

	t = &Types{
		types:  make(map[reflect.Type]uint64),
		coders: make(map[uint64]Type),
		hash:   hash,
	}

	return
}

func (t *Types) GetCoder(typ reflect.Type) (c Type, err error) {
	hash, ok := t.types[typ]

	if !ok {
		err = errors.New("coder not found")
		return
	}

	c, ok = t.coders[hash]

	if !ok {
		err = errors.New("coder not found")
		return
	}

	return
}

func (t *Types) Register(typ reflect.Type) (err error) {
	if _, ok := t.types[typ]; ok {
		err = fmt.Errorf("type %v is already registered", typ)
		return
	}

	t.hash.Reset()

	newT, err := getType(typ, 0, func(k reflect.Kind) {
		t.hash.Write([]byte{byte(k)})
	})

	if err != nil {
		return
	}

	hash := t.hash.Sum64()

	if _, ok := t.coders[hash]; !ok {
		t.coders[hash] = newT
	}

	t.types[typ] = hash

	return
}
