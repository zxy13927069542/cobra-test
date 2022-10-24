package api

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sort"
)

const Path = "/tmp/cobra-test.json"

func loadData() {
	//	打开文件，如果文件不存在则创建
	if file, err := os.OpenFile(Path, os.O_CREATE|os.O_RDWR, 0666); err != nil {
		panic(err)
	} else {
		_ = file.Close()
	}

	var bytes []byte
	var err error
	if bytes, err = ioutil.ReadFile(Path); err != nil {
		panic(err)
	}

	if len(bytes) == 0 {
		return
	}

	if err = json.Unmarshal(bytes, &buffer); err != nil {
		panic(err)
	}
}

func save() {
	if buffer.Data == nil {
		return
	}

	bytes, _ := json.Marshal(buffer)
	if err := ioutil.WriteFile(Path, bytes, 0666); err != nil {
		panic(err)
	}
}

func getLatestIndex() int {
	if buffer.Data == nil || len(buffer.Data) == 0 {
		return -1
	}

	keys := make([]int, 0, 10)
	for k, _ := range buffer.Data {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys[len(keys)-1]
}
