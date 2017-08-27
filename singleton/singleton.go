//http://marcio.io/2015/07/singleton-pattern-in-go/
package singleton

import (
    "sync"
)

type singleton struct {
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
    once.Do(func() {
        instance = &singleton{}
    })
    return instance
}

