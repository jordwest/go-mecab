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

type Model struct {
	model *C.mecab_model_t
}

func (m *Model) getLastError() error {
	cstr := C.mecab_strerror(nil)
	if cstr == nil {
		return nil
	}
	str := C.GoString(cstr)
	if len(str) == 0 {
		return nil
	}

	return errors.New(str)
}

func NewModel(s string) (*Model, error) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))

	model := C.mecab_model_new2(cs)
	m := Model{model: model}
	if model == nil {
		return &m, m.getLastError()
	}

	return &m, nil
}

func (m *Model) Destroy() {
	C.mecab_model_destroy(m.model)
}

func (m *Model) NewTagger() (*Tagger, error) {
	tagger := C.mecab_model_new_tagger(m.model)
	t := Tagger{tagger: tagger}

	if tagger == nil {
		return &t, m.getLastError()
	}

	return &t, nil
}

func (m *Model) NewLattice() (*Lattice, error) {
	lattice := C.mecab_model_new_lattice(m.model)
	l := Lattice{lattice: lattice}

	if lattice == nil {
		return &l, m.getLastError()
	}

	return &l, nil
}
