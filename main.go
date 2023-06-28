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
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("abrsh xo9"))
	})
	errr := http.ListenAndServe(":8080",nil)
	if errr != nil {
		panic(errr.Error())
	}				
}
