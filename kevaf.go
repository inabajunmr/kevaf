/*
Package kevaf is light-weight filebase kvs.
*/
package kevaf

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// NotFoundError express value not exist.
type NotFoundError struct {
	key  string
	path string
}

func (err *NotFoundError) Error() string {
	return fmt.Sprintf("Value is not exist. key:%v, path:%v", err.key, err.path)
}

/*
Map is file base Kvs interface
*/
type Map struct {
	Path string
}

// NewMap is Constructor
func NewMap(path string) (*Map, error) {
	m := new(Map)
	m.Path = path
	testKey := ".writing_test_kevaf"
	err := m.Put(testKey, []byte("test"))
	if err != nil {
		return nil, err
	}

	err = m.Remove(testKey)
	if err != nil {
		return nil, err
	}

	return m, nil
}

/*
Put create file by filename as key and content as value
When failed to write file, return error
*/
func (f Map) Put(key string, value []byte) error {
	file, err := os.OpenFile(createFilePath(f.Path, key),
		os.O_WRONLY|os.O_CREATE, 0666)

	defer file.Close()
	if err != nil {
		return err
	}

	_, err = file.Write(value)
	return err
}

/*
Get read file matching key underneath FileMap.Path
*/
func (f Map) Get(key string) (value []byte, err error) {

	p := createFilePath(f.Path, key)

	if _, err := os.Stat(p); os.IsNotExist(err) {
		return nil, &NotFoundError{key: key, path: f.Path}
	}

	data, err := ioutil.ReadFile(p)
	if err != nil {
		// unexpected error
		return nil, err
	}

	return data, nil
}

/*
Remove specific data by matching key
*/
func (f Map) Remove(key string) (err error) {
	return os.Remove(createFilePath(f.Path, key))
}

/*
RemoveAll data
*/
func (f Map) RemoveAll() (err error) {
	return filepath.Walk(f.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		os.Remove(path)
		return nil
	})
}

func createFilePath(basePath string, key string) string {
	return basePath + "/" + key
}
