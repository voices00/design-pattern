package main

import "fmt"

//	策略模式是一种行为设计模式，
//	它能让你定义一系列算法， 并将每种算法分别放入独立的类中， 以使算法的对象能够相互替换。

// 	思考一下构建内存缓存的情形。 由于处在内存中， 故其大小会存在限制。 在达到其上限后， 一些条目就必须被移除以留出空间。
//	此类操作可通过多种算法进行实现。 一些流行的算法有：
//
//	最少最近使用 （LRU）： 移除最近使用最少的一条条目。
//	先进先出 （FIFO）： 移除最早创建的条目。
//	最少使用 （LFU）： 移除使用频率最低一条条目。
//
//	问题在于如何将我们的缓存类与这些算法解耦， 以便在运行时更改算法。
//	此外， 在添加新算法时， 缓存类不应改变。
//
//	这就是策略模式发挥作用的场景。
//	可创建一系列的算法， 每个算法都有自己的类。 这些类中的每一个都遵循相同的接口， 这使得系列算法之间可以互换。
//	假设通用接口名称为 evictionAlgo移除算法 。
//
//	现在,我们的主要缓存类将嵌入至 evictionAlgo接口中。
//	缓存类会将全部类型的移除算法委派给 evictionAlgo接口， 而不是自行实现。
//	鉴于 evictionAlgo是一个接口， 我们可在运行时将算法更改为 LRU、 FIFO 或者 LFU， 而不需要对缓存类做出任何更改。

func main() {
	lfu := &Lfu{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")

	cache.add("c", "3")

	lru := &Lru{}
	cache.setEvictionAlgo(lru)

	cache.add("d", "4")

	fifo := &Fifo{}
	cache.setEvictionAlgo(fifo)

	cache.add("e", "5")

}

type EvictionAlgo interface {
	evict(c *Cache)
}

type Fifo struct {
}

func (l *Fifo) evict(c *Cache) {
	fmt.Println("Evicting by fifo strtegy")
}

type Lru struct {
}

func (l *Lru) evict(c *Cache) {
	fmt.Println("Evicting by lru strtegy")
}

type Lfu struct {
}

func (l *Lfu) evict(c *Cache) {
	fmt.Println("Evicting by lfu strtegy")
}

type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *Cache) setEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) get(key string) {
	delete(c.storage, key)
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}
