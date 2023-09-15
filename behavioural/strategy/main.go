package strategy

func Main() {
	lfu := &Lfu{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")

	// eviction happens by LFU since cache capacity is max 2
	cache.add("c", "3")

	lru := &Lru{}
	cache.setEvictionAlgo(lru)

	// eviction happens by LRU
	cache.add("d", "4")

	fifo := &Fifo{}
	cache.setEvictionAlgo(fifo)

	// eviction happens by FIFO
	cache.add("e", "5")
}
