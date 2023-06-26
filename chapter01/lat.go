package chapter01

func lat(l []any) bool {
	if len(l) == 0 {
		return true
	} else if atom(car(l)) {
		return lat(cdr(l))
	} else {
		return false
	}
}

func atom(l any) bool {
	if l == nil {
		return false
	}

	switch l.(type) {
	case string:
		return true
	default:
		return false
	}
}

func car(l []any) any {
	return l[0]
}

func cdr(l []any) []any {
	return l[1:]
}
