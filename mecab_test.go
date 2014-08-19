package mecab

import "testing"

func TestVersion(t *testing.T) {
	t.Log("MeCab Version: " + Version())
}

func TestParse(t *testing.T) {
	model, err := NewModel("")
	if err != nil {
		t.Error(err)
	}

	SENTENCE := "This is a test"

	tagger, err := model.NewTagger()
	if err != nil {
		t.Error(err)
	}

	lattice, err := model.NewLattice()
	if err != nil {
		t.Error(err)
	}

	err = lattice.SetSentence(SENTENCE)
	if err != nil {
		t.Errorf("Error setting sentence: %s", err)
	}

	err = tagger.ParseLattice(lattice)
	if err != nil {
		t.Errorf("Error parsing lattice: %s", err)
	}

	t.Logf("Parsed: %s\n\n%s", SENTENCE, lattice)
}
