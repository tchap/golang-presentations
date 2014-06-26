func main() {
	// Uses http.DefaultServeMux.
	http.HandleFunc("/hello", func() (w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "hello, world!\n")
	}
	// Uses http.DefaultServeMux when nil is supplied,
	// which implements http.Handler.
	//
	// Any other http.Handler can be supplied.
	http.ListenAndServe(":8080", nil)
}
