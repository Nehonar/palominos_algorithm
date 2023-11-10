package public

import "io/ioutil"

func LoadIndexFile(filepath string) (string, error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return string(""), err
	}

	return string(content), nil
}
