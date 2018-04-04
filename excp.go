package main

import (
	"os"
	"fmt"
	"strconv"
	"log"
)

func main()  {

	stra:=os.Args[1]
	op:=os.Args[2]
	strb:=os.Args[3]
    a,err:=strconv.Atoi(stra)
    if err!=nil{
    	log.Fatal(err)
	}

	b,err0:=strconv.Atoi(strb)
	if err0!=nil{
		log.Fatal(err0)
	}


if op=="+"{
	fmt.Println(a+b)
}

if op=="*"{
	fmt.Println(a*b)
}

}
