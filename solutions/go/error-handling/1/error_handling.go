package erratum

func Use(opener ResourceOpener, input string) (err error) {
	var resource Resource

	for {
		resource, err = opener()
		if err == nil {
			break
		}
		if _, ok := err.(TransientError); !ok {
			return err
		}
	}

	defer func() {
		resource.Close()
	}()

	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case FrobError:
				// For FrobError, we need to call Defrob before Close
				resource.Defrob(v.defrobTag)
				err = v.inner
			case error:
				// For other errors, just set the error
				err = v
			default:
				// For any other panic type
				panic(r)
			}
		}
	}()

	resource.Frob(input)
	return nil
}
