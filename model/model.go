package model

import (
	"sync"
	"time"
)

type Cache struct {
	// like mysql database
	Base 	map[string]*CacheBase
	Lock 	sync.RWMutex

}

type CacheBase struct {
	// the name of cache database
	Name 		string
	// actually cache
	CacheTable 	map[interface{}]*CacheItem
	Lock		sync.RWMutex

	Add			func()
	Del			func()
	Count		func()


	Timer		func()
}

type CacheItem struct {
	Key 			interface{}
	Value 			interface{}

	CreateTime		time.Time
	UpdateTime		time.Time
	ExpireTime		time.Time
}