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
func TestSave__SavesDataPersistently(t *testing.T){
    t.Parallel()
    path := t.TempDir() + "/kvtest.store"
    s, err := kv.Openstore(path)
    if err != nil {
        t.Fatal(err)
    }
    s.Set("A", "1")
    s.Set("B", "2")
    s.Set("C", "3")
    err = s.Save()
    if err != nil {
        t.Fatal(err)
    }
    s2, err := kv.Openstore(path)
    if err != nil {
        t.Fatal(err)
    }
    if v, _ := s2.Get("A"); v != "1" {
        t.Fatalf("want A=1, got A=%s", v)
    }
    if v, _ := s2.Get("B"); v != "2" {
        t.Fatalf("want B=2, got A=%s", v)
    }
    if v, _ := s2.Get("C"); v != "3" {
        t.Fatalf("want C=3, got A=%s", v)
    }
}