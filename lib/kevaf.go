/*
Package kevaf is light-weight filebase kvs.
*/
package kevaf

import "os"

/*
FileMap is file base Kvs interface
*/
type FileMap struct {
	path string
}

/*
Put create file by filename as key and content as value
When failed to write file, return error
*/
func (f FileMap) Put(key string, value []byte) error {
	file, err := os.Create(f.path)
	defer file.Close()
	if err != nil {
		return err
	}

	file.Write(value)
	return nil
}

func (f FileMap) Get(key string) (value []byte, err error) {
	return nil, nil
}
