package main

import "github.com/ghifar/bookstore-users-api/app"

func main() {
	app.StartApplication()
}

//
//func main() {
//	fmt.Println("Returned:", MyFunc())
//}
//
//func MyFunc() (ret string) {
//	defer func() {
//		if r := recover(); r != nil {
//			ret = "sss"
//		}
//	}()
//	panic("test")
//	return "Normal Return Value"
//}
