
package io

import (
	"reflect"
	"encoding/json"
	"io"
	"io/ioutil"
	"playlist-generator/errors"
)

func UnmarshalGenericFunction(readCloser io.ReadCloser, container interface{}) (responseContainer interface{}, err error) {
	defer readCloser.Close()

	bodyBytes, bytesError := ioutil.ReadAll(readCloser)
	if bytesError != nil {
		return nil, &errors.UnmarshalError{Err: bytesError.Error()}
	}

	typeOfObject := reflect.TypeOf(container)
	responseContainer = reflect.New(typeOfObject)

	err = json.Unmarshal(bodyBytes, &responseContainer)
	if err != nil {
		return nil, &errors.UnmarshalError{Err: err.Error()}
	}

	return responseContainer, nil
}