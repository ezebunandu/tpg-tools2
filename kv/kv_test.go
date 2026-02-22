package kv_test

import (
	"github.com/ezebunandu/kv"
	"testing"
)

func TestGet_ReturnsNotOkIfKeyDoesNotExist(t *testing.T) {
	t.Parallel()
	s, err := kv.OpenStore("dummy path")
	if err != nil {
		t.Fatal(err)
	}
	_, ok := s.Get("key")
	if ok {
		t.Fatal("unexpected ok")
	}
}

func TestGet_ReturnsValueAndOKIfKeyExists(t *testing.T) {
	t.Parallel()
	want := "value"
	s, err := kv.OpenStore("dummy path")
	if err != nil {
		t.Fatal(err)
	}
	s.Set("key", "value")
	got, ok := s.Get("key")
	if !ok {
		t.Fatal("not ok")
	}
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestGet_OverwritesExistingValue(t *testing.T) {
    t.Parallel()
    s, err := kv.OpenStore("dummy path")
    if err != nil {
        t.Fatal(err)
    }
    s.Set("key", "original")
    s.Set("key", "updated")
    v, ok := s.Get("key")
    if !ok{
        t.Fatal("key not found")
    }
    if v != "updated" {
        t.Errorf("want 'updated', got %q", v)
    }
    
}
