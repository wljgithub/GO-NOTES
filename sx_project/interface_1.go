package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache interface {
	Set(key string, value interface{})
	Get(key string) (value interface{}, exist bool)
	Rem(key string)
}

type Icache struct {
	Data     map[string]Value
	Size     int
	Deadline int
	lock     *sync.Mutex
}

type Value struct {
	Elems interface{}
	Times int
}

// 往堆中存入 元素
func (i *Icache) Set(key string, value interface{}) {
	i.lock.Lock()
	defer i.lock.Unlock()

	//如果map超出长度，随机删除一个值
	if len(i.Data) >= i.Size {
		for k := range i.Data {
			delete(i.Data, k)
			break
		}
	}
	// 当map中存在key，不修改value的值
	_, ok := i.Data[key]
	if !ok {
		i.Data[key] = Value{Elems: value, Times: time.Now().YearDay()}
	}
}

// 随机在堆中找一个元素，如果过期则删除
func (i *Icache) Get(key string) (value interface{}, exist bool) {
	i.lock.Lock()
	defer i.lock.Unlock()

	for {
		v, ok := i.Data[key]
		valid := !(v.Times-time.Now().YearDay() > i.Deadline)
		if ok && valid {
			return v.Elems, ok
		} else {
			delete(i.Data, key)
			return
		}
	}
}

// 删除指定元素
func (i *Icache) Rem(key string) {
	i.lock.Lock()
	defer i.lock.Unlock()
	delete(i.Data, key)
}

// 指定每天的 几点 删除堆中所有元素
func (i *Icache) RemOnTime(t int) {
	i.lock.Lock()
	defer i.lock.Unlock()

	if t == time.Now().Hour() {
		for k := range i.Data {
			delete(i.Data, k)
		}
	}
}

func main() {
	//初始化
	i := Icache{
		Size:     10,
		Deadline: 10,
		Data:     map[string]Value{},
		lock:     &sync.Mutex{}}

	// 传入不同类型的值
	i.Set("int", 1)
	i.Set("string", "s")
	i.Set("bool", true)
	i.Set("float64", 9.9)
	for k, v := range i.Data {
		fmt.Println(k, ":", v.Elems)
	}

	//根据key查找值
	v, ok := i.Get("boo")
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("value not exist")
	}

	//	删除指定的值
	i.Rem("int")
	v, ok = i.Get("int")
	if !ok {
		fmt.Println("Value has been remove")
	}

	//在指定时间删除所有的值
	i.RemOnTime(21)
	fmt.Println(i.Data)

}
