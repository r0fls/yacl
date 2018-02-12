package yacl

import (
	"testing"
	"time"
)

var c *cache

func init() {
	c = NewCache()
	c.Insert("key", "value", 1*time.Second)
}

func TestGet(t *testing.T) {
	if c.Get("key") != "value" {
		t.Errorf("Failed test, Got: %v Want: value", c.Get("key"))
	}
}

func TestExpiration(t *testing.T) {
	c.Insert("hey", "dude", 3*time.Second)
	time.Sleep(2 * time.Second)
	if c.Get("key") != nil && c.Get("hey") != "dude" {
		t.Errorf("Failed test, Got: %v Want: value", c.Get("key"))
	}
	time.Sleep(3 * time.Second)
	if c.Get("hey") != nil {
		t.Errorf("Failed test, Got: %v Want: dude", c.Get("key"))
	}
}
