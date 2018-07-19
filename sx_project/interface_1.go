package main

import (
	"fmt"
	// "time"
)

type Cache interface {
	Set(key string, value int)
	Get(key string) (value int, exist bool)
	Rem(key string)
}

type icache struct {
	data map[string]int
	Size int
	// Deadline
}

func (i *icache) Set(key string, value int) {
	//如果map超出长度，随机删除一个值
	if len(i.data) > i.Size {
		for k, _ := range i.data {
			delete(i.data, k)
			break
		}
	}
	// 当map中存在key，不修改value的值
	_, ok := i.data[key]
	if !ok {

		i.data[key] = value
	}
}

func (i *icache) Get(key string) (value int, exist bool) {
	//随机在堆中找一个元素，如果过期则删除

	value, ok := i.data[key]
	if ok {
		return value, ok
	}
	return
}

func (i *icache) Rem(key string) {
	delete(i.data, key)
}

func main() {
	i := icache{map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5}, 5}
	i.Set("6", 6)
	//	fmt.Println(i.data["3"])
	//	fmt.Println(i.Get("4"))
	//	i.Rem("3")
	i.Set("7", 7)
	fmt.Println(i.data)
}
