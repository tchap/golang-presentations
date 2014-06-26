go func() {
	for {
		select {
		case data := <-inputA:
			// foobar
		case data := <-inputB:
			// barfoo
		case <-terminationCh:
			return
		// default:
		}
	}
}
