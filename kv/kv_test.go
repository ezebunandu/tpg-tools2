package kv_test

import (
    "testing"
    "github.com/ezebunandu/kv"
)

func TestGet__ReturnsNotOKIfKeyDoesNotExist(t *testing.T){
    t.Parallel()
    s, err := kv.Openstore("dummy path")
    if err != nil {
        t.Fatal(err)
    }
    _, ok := s.Get("key")
    if ok {
        t.Errorf("unexpected ok")
    }
}

func TestGet__ReturnsValueAndOkIfKeyExists(t *testing.T){
    t.Parallel()
    s, err := kv.Openstore("dummy path")
    if err != nil {
        t.Fatal(err)
    }
    s.Set("key", "value")
    v, ok := s.Get("key")
    if !ok {
        t.Fatal("not ok")
    }
    if v != "value" {
        t.Errorf("want 'value', got %q", v)
    }
}


func TestSet__UpdatesExistingKeyToNewValue(t *testing.T){
    t.Parallel()
    s, err := kv.Openstore("dummy path")
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