package main
//This is a restful with mongodb
import(
	"net/http"
	"fmt"
	"m/router"
	"m/models"
)



func main(){
	//connect db
	models.ConnectDB()

	//routing
	http.HandleFunc("/video", router.TheRouter)
	

	//Listen to server
	fmt.Println("Listining on 9000")
	http.ListenAndServe(":9000", nil)
}

