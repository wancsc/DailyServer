package main

import "sync"

var BucketServer *Bucket

type BucketOptions struct {
	ChannelSize int
}

type Bucket struct {
	lock     sync.RWMutex
	chs      map[string]*Channel
	Operator Operator
}

func NewBucket() (b *Bucket) {
	b = new(Bucket)
	b.chs = make(map[string]*Channel)
	b.Operator = &DefaultOperator{}
	return b
}

func (b *Bucket) Connect(key string, c *Channel) (err error) {
	b.lock.Lock()
	b.chs[key] = c
	//重复判断.
	b.lock.Unlock()
	return
}
func (b *Bucket) Get(key string) (c *Channel) {
	b.lock.RLock()
	c = b.chs[key]
	b.lock.RUnlock()
	return
}

//下线&离线
func (b *Bucket) Offline(Key string) {
	var (
		ok bool
	)
	b.lock.Lock()
	if _, ok = b.chs[Key]; ok {
		delete(b.chs, Key)
	}
	b.lock.Unlock()
}