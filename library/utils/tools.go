package utils

import (
	"github.com/jinzhu/copier"
)

func CopyStruct(src interface{}, dst interface{}) error {
	err := copier.Copy(&dst, &src)
	if err != nil {
		return err
	}
	return nil
}