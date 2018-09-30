package kevaf

import (
	"reflect"
	"testing"

	"github.com/inabajunmr/kevaf/lib"
)

func TestPutGet(t *testing.T) {
	k := "testkey"
	v := []byte{0, 1, 0, 1, 0}

	// initialize
	kvs := kevaf.FileMap{"./testing"}
	err := kvs.Put(k, v)

	if err != nil {
		t.Fatalf("Failed test for Put.", err)
	}

	value, err := kvs.Get(k)
	if err != nil {
		t.Fatalf("Failed test for Put.", err)
	}

	if reflect.DeepEqual(v, value) == false {
		t.Fatalf("Failed test for Get. Not equals between Put value and Get value", err)
	}
}

func TestGetNotExist(t *testing.T) {
	// TODO
}
