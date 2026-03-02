package cacheby_test

import (
	"testing"

	"github.com/sillent/cacheby"
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

func TestCacheBy_SearchStrings(t *testing.T) {
	c := cacheby.NewCacheBy[string, string, string]()
	c.Store("serial", "key1", "value1")
	c.Store("uuid", "key2", "value2")

	val := c.SearchAny("key1")

	if val == nil {
		t.Error("search should find the value")
	} else if *val != "value1" {
		t.Error("search should find correct value")
	}
}

func TestCacheBy_RemoveString(t *testing.T) {
	c := cacheby.NewCacheBy[string, string, string]()
	c.Store("serial", "key1", "value1")
	c.Store("serial", "key2", "value2")

	val := c.Load("serial", "key2")

	if val == nil {
		t.Error("load should return stored value by serial and key")
	} else if *val != "value2" {
		t.Error("loaded value not equal stored by sserial and key")
	}

	c.Remove("serial", "key2")
	val = c.Load("serial", "key2")
	if val != nil {
		t.Error("value should be deleted")
	}
}

func TestCacheBy_LenString(t *testing.T) {
	c := cacheby.NewCacheBy[string, string, string]()
	c.Store("serial", "key1", "value1")
	c.Store("serial", "key2", "value2")
	length := c.Len("serial")
	if length != 2 {
		t.Error("length of value by key name serial should be equal 2")
	}
	length = c.Len("uuid")
	if length != 0 {
		t.Error("length of value by key name uuid should be equal 0")
	}
}

func TestCacheBy_KeysBy(t *testing.T) {
	c := cacheby.NewCacheBy[string, string, string]()
	c.Store("serial", "key1", "value1")
	c.Store("serial", "key2", "value2")
	c.Store("other serial", "key1", "value other")
	keys := c.KeysBy("serial")
	if len(keys) != 2 {
		t.Error("lenght of keys by key name should be equal 2")
	}
}
