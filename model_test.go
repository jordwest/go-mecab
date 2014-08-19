package mecab

import "testing"

func TestNewModel(t *testing.T) {
	m, err := NewModel("")
	defer m.Destroy()

	if err != nil {
		t.Error(err)
	}

	if m == nil {
		t.Error("Model should not be nil")
	}

}

func TestNewTaggerFromModel(t *testing.T) {
	m, err := NewModel("")
	defer m.Destroy()

	if err != nil {
		t.Error(err)
	}
	tagger, err := m.NewTagger()

	if tagger == nil {
		t.Error("Tagger should not be nil")
	}

}
