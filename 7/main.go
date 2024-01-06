package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Cache struct {
	sync.Mutex             //Если только запись, то можно воспользоваться обычным мьютексом
	Data       map[int]int //Есть sync.Map для быстрого чтения данных(проблема cache contention в RWMutex), но у нас идет запись
}

func (c *Cache) Set(key int, value int) {
	c.Lock()
	defer c.Unlock()
	c.Data[key] = value
}

func main() {
	c := Cache{
		Data: make(map[int]int),
	}
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			c.Set(rand.Int(), rand.Int())
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(c.Data)
}
