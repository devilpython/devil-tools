package goroutine_local

import (
	"runtime"
	"sync"
	"fmt"
)

type goroutineLocal struct {
	//initfun func() interface{}
	m       *sync.Map
}

var localDataMap = goroutineLocal{m: &sync.Map{}}

//func NewGoroutineLocal(initfun func() interface{}) *goroutineLocal {
//return &goroutineLocal{initfun: initfun, m: &sync.Map{}}
//}
func GetGoroutineLocal() *goroutineLocal {
	return &localDataMap
}

//获得数据
func (gl *goroutineLocal) Get() interface{} {
	value, ok := gl.m.Load(GetGoroutineID())
	if ok {
		return value
	}else {
		return nil
	}
	//if !ok && gl.initfun != nil {
	//	value = gl.initfun()
	//}
	//return value
}

//设置数据
func (gl *goroutineLocal) Set(v interface{}) {
	gl.m.Store(GetGoroutineID(), v)
}

//删除数据
func (gl *goroutineLocal) Remove() {
	gl.m.Delete(GetGoroutineID())
}

func GetGoroutineID() int {
	var buf [64]byte
	runtime.Stack(buf[:], false)
	var id int
	fmt.Sscanf(string(buf[:]), "goroutine %d", &id)
	return id
}
