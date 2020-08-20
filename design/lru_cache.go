package design

type LRUCache struct {
	Cache    map[int]*Element
	Head     *Element
	Tail     *Element
	Capacity int
}

type Element struct {
	Value int
	Key   int
	Prev  *Element
	Next  *Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cache:    make(map[int]*Element, capacity),
		Capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	e, ok := this.Cache[key]
	if !ok {
		return -1
	}
	if e == this.Head {
		return e.Value
	}
	this.removeElement(e)
	this.MoveToFront(e)
	return e.Value
}

func (this *LRUCache) MoveToFront(e *Element) {
	if e == this.Head {
		return
	}
	// if element is tail
	if e == this.Tail && e.Prev != nil {
		e.Prev.Next = nil
		this.Tail = e.Prev
	}
	e.Prev = nil
	e.Next = this.Head
	if this.Head != nil {
		this.Head.Prev = e
	}
	this.Head = e
}

func (this *LRUCache) Put(key int, value int) {
	e, ok := this.Cache[key]
	if ok {
		e.Value = value
		if e != this.Head {
			this.removeElement(e)
			this.MoveToFront(e)
		}
		return
	}
	if len(this.Cache) >= this.Capacity {
		this.evictElement()
	}
	e = &Element{
		Value: value,
		Key:   key,
	}
	if this.Tail == nil {
		this.Tail = e
	}
	this.Cache[key] = e
	this.MoveToFront(e)
}
func (this *LRUCache) evictElement() {
	if this.Tail == nil {
		return
	}
	key := this.Tail.Key
	// point tail to prev element
	if this.Tail.Prev != nil {
		this.Tail = this.Tail.Prev
		this.Tail.Next = nil
	}
	delete(this.Cache, key)
}
func (this *LRUCache) removeElement(e *Element) {
	// remove element
	if e.Next != nil {
		e.Next.Prev = e.Prev
	}
	if e.Prev != nil {
		e.Prev.Next = e.Next
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
