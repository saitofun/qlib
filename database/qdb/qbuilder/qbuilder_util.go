package qbuilder

func clause(in []byte) []byte {
	return append(append([]byte{'('}, in...), ')')
}

func quoted() {}
