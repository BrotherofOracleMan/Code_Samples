package main

import (
	"time"
	"fmt"
)

func hello(){
	fmt.Println("Hello World go routine")
}

func numbers(){
	for i:=1 ; i <=5 ; i++{
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d",i)
	}
}

func alphabets(){
	for i:= 'a'; i <= 'e'; i++{
		time.Sleep(400)
		fmt.Printf("%c",i)
	}
}

func main(){
	go numbers()
	go alphabets()
	time.Sleep(1* time.Second)
	fmt.Println("main Function")
}