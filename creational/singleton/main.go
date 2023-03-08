package main

import (
	"fmt"
	"sync"
)

//	单例模式是一种创建型设计模式， 让你能够保证一个类只有一个实例， 并提供一个访问该实例的全局节点。

var once sync.Once

type single struct {
}

var singleInstance *single

func getInstance() *single {
	once.Do(
		func() {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		})
	fmt.Println("Single instance already created.")
	return singleInstance
}

func main() {

	for i := 0; i < 10; i++ {
		go getInstance()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
