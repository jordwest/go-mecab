package mecab

/*
#cgo LDFLAGS: -lmecab -lstdc++
#include <mecab.h>
#include <stdlib.h>

*/
import "C"
import "errors"

type Lattice struct {
	lattice *C.mecab_lattice_t
}

func (l *Lattice) getLastError() error {
	cstr := C.mecab_lattice_strerror(l.lattice)
	if cstr == nil {
		return nil
	}
	str := C.GoString(cstr)
	if len(str) == 0 {
		return nil
	}

	return errors.New(str)
}

func (l *Lattice) SetSentence(s string) error {
	cstr := C.CString(s)

	C.mecab_lattice_set_sentence(l.lattice, cstr)

	return l.getLastError()
}

func (l *Lattice) String() string {
	cstr := C.mecab_lattice_tostr(l.lattice)
	return C.GoString(cstr)
}

func (l *Lattice) Destroy() {
	C.mecab_lattice_destroy(l.lattice)
}
