package main

var dbManager *DBManager
var fiberApp *FiberApp

func main() {
	initBackend()
}

func initBackend() {
	//new db manager object
	dbManager = NewDBManager("test")
	dbManager.GetInfo()
	dbManager.setupDb()

	//new fiber app object
	fiberApp = NewFiberApp()
	fiberApp.InitFiber()
}
