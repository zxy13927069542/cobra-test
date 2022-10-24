package api

import (
	"errors"
	"fmt"
)

var buffer Buffer

func Add(name string) {
	loadData()

	if buffer.Data == nil {
		buffer.Data = make(map[int]Person, 10)
	}

	index := getLatestIndex() + 1
	buffer.Data[index] = Person{Id: index, Name: name}
	save()
}

func Del(id int) {
	loadData()

	if buffer.Data == nil {
		return
	}
	delete(buffer.Data, id)
	save()
}

func Get(id int) {
	loadData()

	if buffer.Data == nil {
		panic(errors.New("empty map"))
	}

	if v, exist := buffer.Data[id]; exist {
		fmt.Printf(">>Get() id: [%v] name [%v]\n", id, v.Name)
		return
	}

	panic(errors.New("id not found"))
}

func List() {
	loadData()

	if buffer.Data == nil || len(buffer.Data) == 0 {
		fmt.Println(">>List() ")
		return
	}

	for k, v := range buffer.Data {
		fmt.Printf(">>List() id: [%v] name [%v]\n", k, v.Name)
	}
}
