func (app *App) postPerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	if err := unmarshal(&person, r.Body, 4 * 1024); err != nil { /* boom */ }
	if errs := app.save("people", &person); len(errs) != 0 { /* boom */ }
	w.WriteHeader(http.StatusAccepted)
}

func (app *App) save(collection string, v interface{}) []error {
	if valid, errs := validator.Valid(v); !valid { return errs }
	return []error{app.db.Collection(collection).Insert(v)}
}

func unmarshal(v interface{}, w http.ResponseWriter, r *http.Request, limit int64) error {
	bodyReader := http.MaxBytesReader(w, r.Body, limit)
	defer bodyReader.Close()

	var body bytes.Buffer
	if _, err := io.Copy(body, bodyReader); err != nil { return err }
	if err := json.Unmarshal(body.Bytes(), v); err != nil { return err }
	return nil
}
