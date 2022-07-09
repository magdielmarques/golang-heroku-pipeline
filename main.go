// criando web server
package main

import ( // import de pacotes
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", helloWorld) // criando servidor
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {// atribuindo inicialização do servidor para o servidor web
		panic(err)
	}
}


func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Hello world! It is live now.</h1>")
}
