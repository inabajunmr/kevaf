package kevaf

import (
	"reflect"
	"testing"
)

func TestPutGet(t *testing.T) {
	// TODO user tmp file for test
	k := "testkey"
	v := []byte{0, 1, 0, 1, 0}

	// initialize
	kvs := FileMap{"."}
	err := kvs.Put(k, v)

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

}

func TestGetNotExist(t *testing.T) {
	// TODO
}
