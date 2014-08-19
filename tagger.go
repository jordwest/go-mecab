package mecab

/*
#cgo LDFLAGS: -lmecab -lstdc++
#include <mecab.h>
#include <stdlib.h>

*/
import "C"
import "errors"

type Tagger struct {
	tagger *C.mecab_t
}

func (t *Tagger) getLastError() error {
	cstr := C.mecab_strerror(t.tagger)
	if cstr == nil {
		return nil
	}
	str := C.GoString(cstr)
	if len(str) == 0 {
		return nil
	}

	return errors.New(str)
}

// Clean up the tagger
func (t *Tagger) Destroy() {
	C.mecab_destroy(t.tagger)
}

// Parse a lattice with this tagger
func (t *Tagger) ParseLattice(l *Lattice) error {
	C.mecab_parse_lattice(t.tagger, l.lattice)
	return t.getLastError()
}

/*
TODO

func (t *Tagger) DictionaryInfo() {
	info := C.mecab_dictionary_info(t.tagger)

}
*/
