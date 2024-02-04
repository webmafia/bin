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
	coders map[uint64]coder
	hash   hash.Hash64
}

func NewTypes(key []byte) (t *Types, err error) {
	hash, err := highwayhash.New64(key)

	if err != nil {
		return
	}

	t = &Types{
		types:  make(map[reflect.Type]uint64),
		coders: make(map[uint64]coder),
		hash:   hash,
	}

	return
}

func (t *Types) GetCoder(typ reflect.Type) (c coder, err error) {
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

func (t *Types) Register(typ reflect.Type) (hash uint64, err error) {
	if _, ok := t.types[typ]; ok {
		err = fmt.Errorf("type %v is already registered", typ)
		return
	}

	t.hash.Reset()

	var c coder

	if err = t.dive(&c, typ, 0); err != nil {
		return
	}

	// log.Printf("%+v\n", c.fields)

	hash = t.hash.Sum64()

	if _, ok := t.coders[hash]; !ok {
		t.coders[hash] = c
	}

	t.types[typ] = hash

	return
}

func (t *Types) dive(c *coder, typ reflect.Type, offset uintptr) (err error) {
	switch k := typ.Kind(); k {

	case reflect.Complex64, reflect.Complex128, reflect.Chan, reflect.Func, reflect.Pointer, reflect.Interface, reflect.Map, reflect.UnsafePointer:
		return fmt.Errorf("%s not supported", k)

	case reflect.Slice, reflect.Array:
		t.hashKind(k)
		c.addField(k, offset)

		if err = t.dive(c, typ.Elem(), offset); err != nil {
			return
		}

	case reflect.Struct:
		t.hashKind(k)

		l := typ.NumField()

		for i := 0; i < l; i++ {
			subtyp := typ.Field(i)

			if err = t.dive(c, subtyp.Type, offset+subtyp.Offset); err != nil {
				return
			}
		}

	default:
		t.hashKind(k)
		c.addField(k, offset)
	}

	return
}

func (t *Types) hashKind(kind reflect.Kind) {
	t.hash.Write([]byte{byte(kind)})
}
