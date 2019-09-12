package server

import (
	"testing"
)

func TestNewServer(t *testing.T) {
	_, err := New()
	if err != nil {
		t.Error(err)
	}
}
