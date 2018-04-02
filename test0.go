package main

import (
	"os"
	"strconv"
	"log"
	"fmt"
)

func main()  {
	sum:=0
	for i:=1;i<len(os.Args);i++{
		tmp,err:=strconv.Atoi(os.Args[i])
		if err!=nil{
			log.Fatal(err)
		}
		sum+=tmp
	}
	fmt.Printf("sumary:%d\n",sum)

}