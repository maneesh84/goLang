package maps

import "errors"


type Dictionary map[string]string

func (d Dictionary) Search(key string) (string,error) {
	definition, ok := d[key]

	if(!ok){
		return "", errors.New("word is not in dictionary")
	}

	return definition,nil
}
