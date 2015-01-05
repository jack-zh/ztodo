package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
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
