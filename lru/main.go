package main

//双链表结构


type LRUCache struct {
	size int
	capacity int
	keyMap map[int]*DoubleLinkNode
	head,tail *DoubleLinkNode
}

type DoubleLinkNode struct {
	key,value int
	next,pre *DoubleLinkNode
}

func initDoubleLinkNode(key,value int) *DoubleLinkNode{
	return &DoubleLinkNode{
		key: key,value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l:=LRUCache{
		size: 0,
		capacity: capacity,
		head: initDoubleLinkNode(0,0),
		tail: initDoubleLinkNode(0,0),
		keyMap: make(map[int] *DoubleLinkNode,capacity),
	}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _,ok:=this.keyMap[key];!ok{
		return -1
	}
	node:=this.keyMap[key]
	this.MoveToHead(node)
	return node.value
}


func (this *LRUCache) Put(key int, value int)  {
	if _,ok:=this.keyMap[key];!ok{
		node:=initDoubleLinkNode(key,value)
		this.AddToHead(node)
		this.size++
		if this.size > this.capacity {
			remove:=this.RemoveTail()
			delete(this.keyMap,remove.key)
			this.size--
		}
		this.keyMap[key] = node
	}else{
		node:=this.keyMap[key]
		node.value = value
		this.MoveToHead(node)
	}

}

func (this *LRUCache) RemoveNode(node *DoubleLinkNode){
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) AddToHead(node *DoubleLinkNode){
	node.pre = this.head
	node.next = this.head.next
	this.head.next.pre =node
	this.head.next = node
}

func (this *LRUCache) MoveToHead(node *DoubleLinkNode){
	this.RemoveNode(node)
	this.AddToHead(node)
}

func (this *LRUCache) RemoveTail()*DoubleLinkNode{
	node:=this.tail.pre
	this.RemoveNode(node)
	return node
}
