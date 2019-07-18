package main

import (
	"fmt"
	"math"
	"net/http"
)

func greeting(text string) string {
	return fmt.Sprintf("<b>%s</b>", text)
}

func loop() float64 {
	x := 0.0001
	for index := 0; index < 100000000; index++ {
		x = x + math.Sqrt(x)
	}
	return x
}

func main() {
 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		loop()
		fmt.Fprintf(w, greeting("Code.education Rocks!"))
 	})
 	http.ListenAndServe(":80", nil)
}