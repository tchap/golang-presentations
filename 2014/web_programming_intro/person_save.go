type Person struct {
	app   *App // HL
	Name  string `validator:"max=64,nonzero"`
	Phone int    `validator:"regexp=..."     bson:"PhOnE"`
}

func (p *Person) Save() error {
	return p.app.save("people", p) // HL
}

func (app *App) savePerson(w http.ResponseWriter, r *http.Request) {
	person := &Person{app: app}
	if err := unmarshal(person, r.Body, 4 * 1024); err != nil { /* boom */ }
	if err := person.Save(); err != nil { /* boom */ }
	w.WriteHeader(http.StatusAccepted)
}
