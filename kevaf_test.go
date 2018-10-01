package kevaf

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestPutGet(t *testing.T) {
	k := "testkey"
	v := []byte("hello")

	dir, err := ioutil.TempDir("", "kevaf_test")
	if err != nil {
		t.Fatal("Can not prepare test.", err)
	}

	// initialize
	kvs := Map{dir}
	err = kvs.Put(k, v)

	if err != nil {
		t.Fatal("Failed test for Put.", err)
	}

	value, err := kvs.Get(k)

	if err != nil {
		t.Fatal("Failed test for Put.", err)
	}

	if reflect.DeepEqual(v, value) == false {
		t.Fatal("Failed test for Get. Not equals between Put value and Get value", err)
	}

	// update
	v = []byte("goodbye")
	err = kvs.Put(k, v)

	if err != nil {
		t.Fatal("Failed test for Put.", err)
	}

	value, err = kvs.Get(k)

	if err != nil {
		t.Fatal("Failed test for Put.", err)
	}

	if reflect.DeepEqual(v, value) == false {
		t.Fatal("Failed test for Get. Not equals between Put value and Get value", err)
	}
}

func TestGetNotExist(t *testing.T) {
	dir, err := ioutil.TempDir("", "kevaf_test")
	if err != nil {
		t.Fatal("Can not prepare test.", err)
	}

	kvs := Map{dir}
	_, err = kvs.Get("absent")

	if err == nil {
		t.Fatal("Not put but err is nil.")
	}
}
