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
	kvs, err := NewMap(dir)
	if err != nil {
		t.Fatal("Failed to initialize.", err)
	}
	err = kvs.Put(k, v)

	if err != nil {
		t.Fatal("Failed test for Put.", err)
	}

	value, err := kvs.Get(k)

	if err != nil {
		t.Fatal("Failed test for Put.", err)
	}

	if !reflect.DeepEqual(v, value) {
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

	if !reflect.DeepEqual(v, value) {
		t.Fatal("Failed test for Get. Not equals between Put value and Get value", err)
	}
}

func TestGetNotExist(t *testing.T) {
	dir, err := ioutil.TempDir("", "kevaf_test")
	if err != nil {
		t.Fatal("Can not prepare test.", err)
	}

	kvs, err := NewMap(dir)
	if err != nil {
		t.Fatal("Failed to initialize.", err)
	}
	_, err = kvs.Get("absent")

	typeName := reflect.TypeOf(err).String()
	expectedTypeName := "*kevaf.NotFoundError"
	if typeName != expectedTypeName {
		t.Fatalf("Error type is unexpected. Expected:%v, Actual:%v", expectedTypeName, typeName)
	}

	if err == nil {
		t.Fatal("Not put but err is nil.")
	}
}

func TestRemove(t *testing.T) {
	k := "testkey"
	v1 := []byte("hello")

	dir, err := ioutil.TempDir("", "kevaf_test")
	if err != nil {
		t.Fatal("Can not prepare test.", err)
	}

	// initialize
	kvs, err := NewMap(dir)
	if err != nil {
		t.Fatal("Failed to initialize.", err)
	}
	err = kvs.Put(k, v1)

	if err != nil {
		t.Fatal("Failed test for Put.", err)
	}

	_, err = kvs.Get(k)
	if err != nil {
		t.Fatal("Failed test for Put.", err)
	}

	// Remove
	err = kvs.Remove(k)

	if err != nil {
		t.Fatal("Failed test for Put.", err)
	}

	_, err = kvs.Get(k)
	if err == nil {
		t.Fatal("Get removed value but no error.", err)
	}

}

func TestRemoveAll(t *testing.T) {
	k1 := "testkey1"
	v1 := []byte("hello")

	dir, err := ioutil.TempDir("", "kevaf_test")
	if err != nil {
		t.Fatal("Can not prepare test.", err)
	}

	// initialize
	kvs, err := NewMap(dir)
	if err != nil {
		t.Fatal("Failed to initialize.", err)
	}

	err = kvs.Put(k1, v1)

	if err != nil {
		t.Fatal("Failed test for Put.", err)
	}

	k2 := "testkey2"
	v2 := []byte("goodbye")

	err = kvs.Put(k2, v2)

	if err != nil {
		t.Fatal("Failed test for Put.", err)
	}

	kvs.RemoveAll()

	_, err = kvs.Get(k1)
	if err == nil {
		t.Fatal("RemoveAll but value returned")
	}

	_, err = kvs.Get(k2)
	if err == nil {
		t.Fatal("RemoveAll but value returned")
	}

}

func TestNewMapNotExistDir(t *testing.T) {
	_, err := NewMap("notexistdir")
	if err == nil {
		t.Fatal("Not returned expected error.")
	}
}
