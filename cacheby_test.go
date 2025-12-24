package cacheby_test

import (
	"github.com/sillent/cacheby"
	"testing"
)

func TestCacheBy_StoreStrings(t *testing.T) {
	c := cacheby.NewCacheBy[string, string, string]()
	c.Store("serial", "key1", "value1")
	val := c.Load("serial", "key1")

	if val == nil {
		t.Error("load should return stored value by serial and key")
	} else if *val != "value1" {
		t.Error("loaded value not equal stored by serial and key")
	}

	c.Store("uuid", "key1", "value2")
	val = c.Load("uuid", "key1")

	if val == nil {
		t.Error("load should return stored value by uuid and key")
	} else if *val != "value2" {
		t.Error("loaded value not equal stored by uuid and key")
	}

}
func TestCacheBy_StoreInteger(t *testing.T) {
	c := cacheby.NewCacheBy[string, string, int]()
	c.Store("serial", "key1", 1)
	val := c.Load("serial", "key1")

	if val == nil {
		t.Error("load should return stored value by serial and key")
	} else if *val != 1 {
		t.Error("loaded value not equal stored by serial and key")
	}

	c.Store("uuid", "key1", 2)
	val = c.Load("uuid", "key1")

	if val == nil {
		t.Error("load should return stored value by uuid and key")
	} else if *val != 2 {
		t.Error("loaded value not equal stored by uuid and key")
	}

}
