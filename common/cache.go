package common

import "sync"

type Cache struct {
	Tenamts		map[string]*Tenant
	Lock  		sync.RWMutex
}


func InitCache() {

}