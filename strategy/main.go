package main

func main() {
	lfu := &Lfu{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")

	cache.add("c", "3")

	lru := &Lru{}

	cache.setEvictionAlog(lru)

	cache.add("d", "4")

	fifo := &Fifo{}
	cache.setEvictionAlog(fifo)

	cache.add("e", "5")
}
