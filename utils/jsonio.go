package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

func Str2Map(jsonStr string) (map[string]interface{}, error) {
	var jsonObj map[string]interface{}

	if err := json.Unmarshal([]byte(jsonStr), &jsonObj); err != nil {
		return nil, err
	}
	return jsonObj, nil
}

func ReadJsonFile2Map(filename string) (map[string]interface{}, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return Str2Map(string(bytes))
}

func Map2File(mapobj map[string]interface{}, filename string) error {
	fd, e := os.Create(filename)
	defer fd.Close()
	if e != nil {
		return e
	}
	enc := json.NewEncoder(fd)
	enc.Encode(mapobj)
	return nil
}

func Struct2Map(obj interface{}) (map[string]interface{}, error) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		fmt.Println("...")
		key := t.Field(i).Name
		value := v.FieldByName(key).String()
		data[key] = value
	}
	return data, nil
}

func Struct2File(obj interface{}) error {
	mapObj, err := Struct2Map(obj)
	if err != nil {
		return err
	}
	return Map2File(mapObj, "nil")
}
