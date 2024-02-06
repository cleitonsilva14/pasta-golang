package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
)

func main(){
	
	//stdLog = log.New(os.Stdout, "INFO", 0)
	//stdLog.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(os.Stdout)
	
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/login", loginPage)

	log.Printf("Server running: http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}

func rootHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Hello World</h1><div>Welcome to Golang WebApp</div>")
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	text := `
		<h1 style="">Login Page</h1>
		<div>
			<form>
				<div>
					<label for="inputUser">User:</label>
					<input type="text" id="inputUser">
				</div>
				</div>
					<label for="inputPassword">Password:</label>
					<input type="password" id="inputPassword">
				</div>
			</form>
		</div>
	` 
	fmt.Fprintf(w, text)
}

