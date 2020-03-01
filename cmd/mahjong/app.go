package main

import "sync"

var wg sync.WaitGroup

//StopServer 当前不需要做什么
func StopServer(){

}

func Run(f func()) {
	wg.Add(1)
	defer wg.Done()
	f()
}

func StartServer(){
	Run(Gin)
	Run(NaNo)
}

func Gin() {

}

func NaNo() {

}
