package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func runBash(w http.ResponseWriter, r *http.Request) {
	if _, err := os.Stat(".phteven.sh"); os.IsNotExist(err) {
		const fileDoesNotExist = "File `.phteven.sh` does not exist"
		log.Print(fileDoesNotExist)
		fmt.Fprintf(w, fileDoesNotExist)
		return
	}

	out, err := exec.Command("/bin/sh", ".phteven.sh").Output()
	outString := string(out)
	if err != nil {
		log.Printf("[1] %s", outString)
		http.Error(w, outString, http.StatusInternalServerError)
		return
	}

	log.Printf("[0] %s", outString)
	fmt.Fprintf(w, outString)
}

func main() {
	port := "9847"
	if len(os.Args) >= 2 {
		port = os.Args[1]
	}
	portWithColon := fmt.Sprintf(":%s", port)

	http.HandleFunc("/", runBash)
	log.Printf("Phteven running. Serving on http://0.0.0.0%s ... ", portWithColon)
	http.ListenAndServe(portWithColon, nil)
}
