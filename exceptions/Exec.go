package exceptions

func Exec0r0[E error](f func() E) {
	err := f()
	if error(err) != nil {
		panic(Package(err))
	}
}

func Exec0r1[R1 any, E error](f func() (R1, E)) R1 {
	r, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r
}

func Exec0r2[R1, R2 any, E error](f func() (R1, R2, E)) (R1, R2) {
	r1, r2, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec1r0[A1 any, E error](f func(A1) E, a1 A1) {
	err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
}

func Exec1r1[A1, R1 any, E error](f func(A1) (R1, E), a1 A1) R1 {
	r, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r
}

func Exec1r2[A1, R1, R2 any, E error](f func(A1) (R1, R2, E), a1 A1) (R1, R2) {
	r1, r2, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec2r0[A1, A2 any, E error](f func(A1, A2) E, a1 A1, a2 A2) {
	err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
}

func Exec2r1[A1, A2, R1 any, E error](f func(A1, A2) (R1, E), a1 A1, a2 A2) R1 {
	r, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r
}

func Exec2r2[A1, A2, R1, R2 any, E error](f func(A1, A2) (R1, R2, E), a1 A1, a2 A2) (R1, R2) {
	r1, r2, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec3r0[A1, A2, A3 any, E error](f func(A1, A2, A3) E, a1 A1, a2 A2, a3 A3) {
	err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
}

func Exec3r1[A1, A2, A3, R1 any, E error](f func(A1, A2, A3) (R1, E), a1 A1, a2 A2, a3 A3) R1 {
	r, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r
}

func Exec3r2[A1, A2, A3, R1, R2 any, E error](f func(A1, A2, A3) (R1, R2, E), a1 A1, a2 A2, a3 A3) (R1, R2) {
	r1, r2, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}
