/*
Package kevaf is light-weight filebase kvs.
*/
package kevaf

import (
	"io/ioutil"
	"os"
)

/*
Map is file base Kvs interface
*/
type Map struct {
	Path string
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
	data, err := ioutil.ReadFile(createFilePath(f.Path, key))
	if err != nil {
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

func createFilePath(basePath string, key string) string {
	return basePath + "/" + key
}
