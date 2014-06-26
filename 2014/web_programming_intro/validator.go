type Person struct {
	Name  string `validator:"max=64,nonzero"`
	Phone int    `validator:"regexp=..."     bson:"PhOnE"`
}

p := &Person{Name: "Pepa Novak"}
if valid, errs := validator.Validate(p); !valid {
	// Return an error
}
