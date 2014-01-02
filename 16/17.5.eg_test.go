package eg

import (
	"testing"
)

func TestExample(t *testing.T) {
	e := NewExample()

	_ = e.(*concreteType)

	if e.Name() != "Concrete Type" {
		t.Fail()
	}

	t.Errorf("This test is buggy")
}
