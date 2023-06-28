package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
)
func main(){
	res,err := http.Get("https://catfact.ninja/fact")
	if err != nil{
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots,err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(robots))
}