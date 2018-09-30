/*
Package kevaf is light-weight filebase kvs.
*/
package kevaf

import (
	"io/ioutil"
	"os"
)

/*
FileMap is file base Kvs interface
*/
type FileMap struct {
	Path string
}

/*
Put create file by filename as key and content as value
When failed to write file, return error
*/
func (f FileMap) Put(key string, value []byte) error {
	file, err := os.OpenFile(f.Path+"/"+key, os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		return err
	}

	file.Write(value)
	return nil
}

/*
Get read file matching key underneath FileMap.Path
*/
func (f FileMap) Get(key string) (value []byte, err error) {
	data, err := ioutil.ReadFile(f.Path + "/" + key)
	if err != nil {
		return nil, err
	}

	return data, nil
}
