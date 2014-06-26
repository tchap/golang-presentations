func (app *App) save(v interface{}) error {
	typeName := reflect.Indirect(reflect.ValueOf(v)).Type().Name()
	if typeName == "" {
		return errors.New("unexported type cannot be saved")
	}
	typeName = strings.ToLower(typeName)
	return app.db.Collection(typeName).Insert(v)
}

func (app *App) validate(v interface{}) (valid bool, errs []error) {
	return validator.Valid(v)
}
