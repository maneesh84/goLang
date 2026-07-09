package maps

import "errors"

var ErrorNotFound = errors.New("word not found in dictionary")

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string,error) {
	definition, ok := d[key]

	if(!ok){
		return "", ErrorNotFound
	}

	return definition,nil
}
