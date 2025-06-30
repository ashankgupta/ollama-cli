package cmd

import (
	"encoding/base64"
	"io/ioutil"
)

func EncodeImageToBase64(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

