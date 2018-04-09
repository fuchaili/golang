package main

import (
	"os"
	"log"
	"fmt"
	"strings"
	"net/http"

	"io/ioutil"
)

func  printFiles(name string) {
	resp, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(string(resp))
		fmt.Println()
	}
}

func printHTTP(url string)  {
resp,err:=http.Get(url)
if err!=nil{
	fmt.Println(err)
}else {
	fmt.Println(resp)
}


}



func main()  {
	var name string
	name=os.Args[1]
	fmt.Println(name)
	log.Print(name)
     if strings.HasPrefix(name,"http://"){
         printHTTP(name)
	 }else {
	 	printFiles(name)
	 }

}
