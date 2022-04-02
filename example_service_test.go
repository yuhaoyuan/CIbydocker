package utbydocker

import "testing"

func TestExampleUT(t *testing.T) {
	err := NewRedis()
	if err != nil {
		t.Error(err)
	}

	err = NewGORM()
	if err != nil {
		t.Error(err)
	}
}
