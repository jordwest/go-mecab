package mecab

/*
#cgo LDFLAGS: -lmecab -lstdc++
#include <mecab.h>
#include <stdlib.h>

*/
import "C"
import (
	"errors"
	"unsafe"
)

var ErrNoMoreTokens = errors.New("No more tokens")

type Iterator interface {
	Next() (string, error)
}

type Tokenizer interface {
	Tokenize(string) Iterator
}

type TokenizeMecab struct {
	mecab *C.mecab_t
}

type TokenizeMecabIter struct {
	root    *C.mecab_node_t
	current *C.struct_mecab_node_t
}

func NewMecab(s string) (*TokenizeMecab, error) {
	config := C.CString(s)
	newMecabPtr := C.mecab_new2(config)
	if newMecabPtr == nil {
		errStrPtr := C.mecab_strerror(newMecabPtr)
		return nil, errors.New(C.GoString(errStrPtr))
	}
	return &TokenizeMecab{mecab: newMecabPtr}, nil
}

func (t *TokenizeMecab) Tokenize(input string) *TokenizeMecabIter {
	p := C.CString(input)
	defer C.free(unsafe.Pointer(p))

	node := C.mecab_sparse_tonode(t.mecab, p)
	return &TokenizeMecabIter{node, node.next}
}

func (iter *TokenizeMecabIter) Next() (string, error) {
	if iter.current == nil {
		return "", ErrNoMoreTokens
	}

	node := iter.current
	iter.current = iter.current.next

	s := C.GoString(node.surface)
	return s[:int(node.length)], nil
}

func Version() string {
	cstr := C.mecab_version()
	return C.GoString(cstr)
}
