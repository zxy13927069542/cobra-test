package api

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Buffer struct {
	Data map[int]Person `json:"data"`
}
