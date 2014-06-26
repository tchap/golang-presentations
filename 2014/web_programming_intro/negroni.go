func main() {
	// Middleware after the mux and before HelloServer, i.e. local to this path.
	pathChain := negroni.New(middleware1, middleware2, middleware3)
	pathChain.UseHandler(HelloHandler)

	mux := pat.New()
	mux.Get("/hello/:name", pathChain)

	// Middleware before the mux, i.e. shared for all requests.
	sharedChain := negroni.New(middlewareA, middlewareB, middlewareC)
	sharedChain.UseHandler(mux)

	// web --> sharedChain --> mux --> pathChain
	http.Handle("/", sharedChain)
	if err := http.ListenAndServe(":12345", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
