package main
import
(
	"fmt"
	"net/http"
	"errors"
	"os"
	"io"
)


func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func getHello(w http.ResponseWriter, r *http.Request){
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "This is my website!\n")


}


func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)


		err:= http.ListenAndServe(":3333", nil)
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Server closed")
		} else if err != nil {
			fmt.Printf("error starting server: %s\n", err)
			os.Exit(1)
		}
}