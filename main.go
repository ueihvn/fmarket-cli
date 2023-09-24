package main

import (
	"github.com/ueihvn/fmarket-cli/cmd"
)

func main() {
	cmd.Execute()
	/*
		http.HandleFunc("/favicon.ico", func(res http.ResponseWriter, req *http.Request) {
			res.Write([]byte{})
		})
		http.HandleFunc("/custom", drawCustomChart)
		http.ListenAndServe(":8080", nil)
	*/
}
