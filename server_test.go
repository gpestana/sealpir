package sealpir

import (
	"testing"
)

func TestNewServer(t *testing.T) {
	_, err := NewServer()
	if err != nil {
		t.Error(err)
	}
}
