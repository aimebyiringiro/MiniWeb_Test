package main

// I am trying to change something
import (
	"addVWeb/controllers"
	"net/http"
)

func main() {

	controllers.RegesterController()
	http.ListenAndServe(" :3000", nil)

	// the IP address were are listing to
	// the objec that is to handle all the request are coming
	// What is the front controller and what is the back controller

}
