func app.savePerson(w http.ResponseWriter, r *http.Request) {
	bodyReader := http.MaxBytesReader(w, r.Body, 4 * 1024)
	defer bodyReader.Close()

	var body bytes.Buffer
	if _, err := io.Copy(bodyReader, body); err != nil {
		http.Error(w, "Person Too Large", http.StatusRequestEntityTooLarge)
		return
	}

	var person Person
	if err := json.Unmarshal(body.Bytes(), &person); err != nil { /* boom */ }

	if valid, errs := validator.Valid(&person); !valid { /* boom */ }

	if err := app.db.Collection("people").Insert(&person); err != nil { /* boom */ }

	w.WriteHeader(http.StatusAccepted)
}
