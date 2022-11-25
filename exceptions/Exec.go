/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package exceptions

func Exec0r0[E error](f func() E) {
	err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec0r1[R1 any, E error](f func() (R1, E)) R1 {
	r1, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec0r2[R1, R2 any, E error](f func() (R1, R2, E)) (R1, R2) {
	r1, r2, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec0r3[R1, R2, R3 any, E error](f func() (R1, R2, R3, E)) (R1, R2, R3) {
	r1, r2, r3, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec0r4[R1, R2, R3, R4 any, E error](f func() (R1, R2, R3, R4, E)) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec0r5[R1, R2, R3, R4, R5 any, E error](f func() (R1, R2, R3, R4, R5, E)) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec0r6[R1, R2, R3, R4, R5, R6 any, E error](f func() (R1, R2, R3, R4, R5, R6, E)) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec0r7[R1, R2, R3, R4, R5, R6, R7 any, E error](f func() (R1, R2, R3, R4, R5, R6, R7, E)) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec0r8[R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func() (R1, R2, R3, R4, R5, R6, R7, R8, E)) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec0r9[R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func() (R1, R2, R3, R4, R5, R6, R7, R8, R9, E)) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec0r10[R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func() (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E)) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec0r11[R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func() (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E)) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec0r12[R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func() (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E)) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec0r13[R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func() (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E)) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec0r14[R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func() (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E)) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec0r15[R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func() (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E)) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec0r16[R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func() (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E)) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec0r17[R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func() (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E)) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec0r18[R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func() (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E)) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec0r19[R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func() (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E)) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec0r20[R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func() (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E)) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec0r21[R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func() (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E)) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec0r22[R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func() (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E)) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec0r23[R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func() (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E)) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec0r24[R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func() (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E)) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec0r25[R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func() (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E)) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec0r26[R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func() (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E)) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec0r27[R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func() (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E)) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec0r28[R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func() (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E)) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec0r29[R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func() (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E)) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec0r30[R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func() (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E)) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec0r31[R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func() (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E)) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f()
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec1r0[A1 any, E error](f func(A1) E, a1 A1) {
	err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec1r1[A1, R1 any, E error](f func(A1) (R1, E), a1 A1) R1 {
	r1, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec1r2[A1, R1, R2 any, E error](f func(A1) (R1, R2, E), a1 A1) (R1, R2) {
	r1, r2, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec1r3[A1, R1, R2, R3 any, E error](f func(A1) (R1, R2, R3, E), a1 A1) (R1, R2, R3) {
	r1, r2, r3, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec1r4[A1, R1, R2, R3, R4 any, E error](f func(A1) (R1, R2, R3, R4, E), a1 A1) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec1r5[A1, R1, R2, R3, R4, R5 any, E error](f func(A1) (R1, R2, R3, R4, R5, E), a1 A1) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec1r6[A1, R1, R2, R3, R4, R5, R6 any, E error](f func(A1) (R1, R2, R3, R4, R5, R6, E), a1 A1) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec1r7[A1, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec1r8[A1, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec1r9[A1, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec1r10[A1, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec1r11[A1, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec1r12[A1, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec1r13[A1, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec1r14[A1, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec1r15[A1, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec1r16[A1, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec1r17[A1, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec1r18[A1, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec1r19[A1, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec1r20[A1, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec1r21[A1, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec1r22[A1, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec1r23[A1, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec1r24[A1, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec1r25[A1, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec1r26[A1, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec1r27[A1, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec1r28[A1, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec1r29[A1, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec1r30[A1, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec1r31[A1, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec2r0[A1, A2 any, E error](f func(A1, A2) E, a1 A1, a2 A2) {
	err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec2r1[A1, A2, R1 any, E error](f func(A1, A2) (R1, E), a1 A1, a2 A2) R1 {
	r1, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec2r2[A1, A2, R1, R2 any, E error](f func(A1, A2) (R1, R2, E), a1 A1, a2 A2) (R1, R2) {
	r1, r2, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec2r3[A1, A2, R1, R2, R3 any, E error](f func(A1, A2) (R1, R2, R3, E), a1 A1, a2 A2) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec2r4[A1, A2, R1, R2, R3, R4 any, E error](f func(A1, A2) (R1, R2, R3, R4, E), a1 A1, a2 A2) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec2r5[A1, A2, R1, R2, R3, R4, R5 any, E error](f func(A1, A2) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec2r6[A1, A2, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec2r7[A1, A2, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec2r8[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec2r9[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec2r10[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec2r11[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec2r12[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec2r13[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec2r14[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec2r15[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec2r16[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec2r17[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec2r18[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec2r19[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec2r20[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec2r21[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec2r22[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec2r23[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec2r24[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec2r25[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec2r26[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec2r27[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec2r28[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec2r29[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec2r30[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec2r31[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec3r0[A1, A2, A3 any, E error](f func(A1, A2, A3) E, a1 A1, a2 A2, a3 A3) {
	err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec3r1[A1, A2, A3, R1 any, E error](f func(A1, A2, A3) (R1, E), a1 A1, a2 A2, a3 A3) R1 {
	r1, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec3r2[A1, A2, A3, R1, R2 any, E error](f func(A1, A2, A3) (R1, R2, E), a1 A1, a2 A2, a3 A3) (R1, R2) {
	r1, r2, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec3r3[A1, A2, A3, R1, R2, R3 any, E error](f func(A1, A2, A3) (R1, R2, R3, E), a1 A1, a2 A2, a3 A3) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec3r4[A1, A2, A3, R1, R2, R3, R4 any, E error](f func(A1, A2, A3) (R1, R2, R3, R4, E), a1 A1, a2 A2, a3 A3) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec3r5[A1, A2, A3, R1, R2, R3, R4, R5 any, E error](f func(A1, A2, A3) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2, a3 A3) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec3r6[A1, A2, A3, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2, A3) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2, a3 A3) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec3r7[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2, a3 A3) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec3r8[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2, a3 A3) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec3r9[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2, a3 A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec3r10[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2, a3 A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec3r11[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2, a3 A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec3r12[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2, a3 A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec3r13[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2, a3 A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec3r14[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2, a3 A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec3r15[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2, a3 A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec3r16[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2, a3 A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec3r17[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2, a3 A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec3r18[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2, a3 A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec3r19[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2, a3 A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec3r20[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2, a3 A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec3r21[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2, a3 A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec3r22[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2, a3 A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec3r23[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2, a3 A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec3r24[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2, a3 A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec3r25[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2, a3 A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec3r26[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2, a3 A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec3r27[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2, a3 A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec3r28[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2, a3 A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec3r29[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2, a3 A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec3r30[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2, a3 A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec3r31[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2, a3 A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2, a3)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec4r0[A1, A2, A3, A4 any, E error](f func(A1, A2, A3, A4) E, a1 A1, a2 A2, a3 A3, a4 A4) {
	err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec4r1[A1, A2, A3, A4, R1 any, E error](f func(A1, A2, A3, A4) (R1, E), a1 A1, a2 A2, a3 A3, a4 A4) R1 {
	r1, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec4r2[A1, A2, A3, A4, R1, R2 any, E error](f func(A1, A2, A3, A4) (R1, R2, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec4r3[A1, A2, A3, A4, R1, R2, R3 any, E error](f func(A1, A2, A3, A4) (R1, R2, R3, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec4r4[A1, A2, A3, A4, R1, R2, R3, R4 any, E error](f func(A1, A2, A3, A4) (R1, R2, R3, R4, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec4r5[A1, A2, A3, A4, R1, R2, R3, R4, R5 any, E error](f func(A1, A2, A3, A4) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec4r6[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec4r7[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec4r8[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec4r9[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec4r10[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec4r11[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec4r12[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec4r13[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec4r14[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec4r15[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec4r16[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec4r17[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec4r18[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec4r19[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec4r20[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec4r21[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec4r22[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec4r23[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec4r24[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec4r25[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec4r26[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec4r27[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec4r28[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec4r29[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec4r30[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec4r31[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2, a3, a4)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec5r0[A1, A2, A3, A4, A5 any, E error](f func(A1, A2, A3, A4, A5) E, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) {
	err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec5r1[A1, A2, A3, A4, A5, R1 any, E error](f func(A1, A2, A3, A4, A5) (R1, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R1 {
	r1, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec5r2[A1, A2, A3, A4, A5, R1, R2 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec5r3[A1, A2, A3, A4, A5, R1, R2, R3 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, R3, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec5r4[A1, A2, A3, A4, A5, R1, R2, R3, R4 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, R3, R4, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec5r5[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec5r6[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec5r7[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec5r8[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec5r9[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec5r10[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec5r11[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec5r12[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec5r13[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec5r14[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec5r15[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec5r16[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec5r17[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec5r18[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec5r19[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec5r20[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec5r21[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec5r22[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec5r23[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec5r24[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec5r25[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec5r26[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec5r27[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec5r28[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec5r29[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec5r30[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec5r31[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2, a3, a4, a5)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec6r0[A1, A2, A3, A4, A5, A6 any, E error](f func(A1, A2, A3, A4, A5, A6) E, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) {
	err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec6r1[A1, A2, A3, A4, A5, A6, R1 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R1 {
	r1, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec6r2[A1, A2, A3, A4, A5, A6, R1, R2 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec6r3[A1, A2, A3, A4, A5, A6, R1, R2, R3 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, R3, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec6r4[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec6r5[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec6r6[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec6r7[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec6r8[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec6r9[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec6r10[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec6r11[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec6r12[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec6r13[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec6r14[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec6r15[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec6r16[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec6r17[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec6r18[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec6r19[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec6r20[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec6r21[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec6r22[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec6r23[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec6r24[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec6r25[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec6r26[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec6r27[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec6r28[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec6r29[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec6r30[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec6r31[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2, a3, a4, a5, a6)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec7r0[A1, A2, A3, A4, A5, A6, A7 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) E, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) {
	err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec7r1[A1, A2, A3, A4, A5, A6, A7, R1 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R1 {
	r1, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec7r2[A1, A2, A3, A4, A5, A6, A7, R1, R2 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec7r3[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec7r4[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec7r5[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec7r6[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec7r7[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec7r8[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec7r9[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec7r10[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec7r11[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec7r12[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec7r13[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec7r14[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec7r15[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec7r16[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec7r17[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec7r18[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec7r19[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec7r20[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec7r21[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec7r22[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec7r23[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec7r24[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec7r25[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec7r26[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec7r27[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec7r28[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec7r29[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec7r30[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec7r31[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2, a3, a4, a5, a6, a7)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec8r0[A1, A2, A3, A4, A5, A6, A7, A8 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) E, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) {
	err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec8r1[A1, A2, A3, A4, A5, A6, A7, A8, R1 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R1 {
	r1, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec8r2[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec8r3[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec8r4[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec8r5[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec8r6[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec8r7[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec8r8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec8r9[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec8r10[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec8r11[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec8r12[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec8r13[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec8r14[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec8r15[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec8r16[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec8r17[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec8r18[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec8r19[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec8r20[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec8r21[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec8r22[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec8r23[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec8r24[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec8r25[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec8r26[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec8r27[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec8r28[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec8r29[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec8r30[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec8r31[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2, a3, a4, a5, a6, a7, a8)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec9r0[A1, A2, A3, A4, A5, A6, A7, A8, A9 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) E, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) {
	err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec9r1[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R1 {
	r1, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec9r2[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec9r3[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec9r4[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec9r5[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec9r6[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec9r7[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec9r8[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec9r9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec9r10[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec9r11[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec9r12[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec9r13[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec9r14[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec9r15[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec9r16[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec9r17[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec9r18[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec9r19[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec9r20[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec9r21[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec9r22[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec9r23[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec9r24[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec9r25[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec9r26[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec9r27[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec9r28[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec9r29[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec9r30[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec9r31[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec10r0[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) E, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) {
	err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec10r1[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) R1 {
	r1, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec10r2[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec10r3[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2, R3 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, R3, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec10r4[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2, R3, R4 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, R3, R4, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec10r5[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2, R3, R4, R5 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec10r6[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec10r7[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec10r8[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec10r9[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec10r10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec10r11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec10r12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec10r13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec10r14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec10r15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec10r16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec10r17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec10r18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec10r19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec10r20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec10r21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec10r22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec10r23[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec10r24[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec10r25[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec10r26[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec10r27[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec10r28[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec10r29[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec10r30[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec10r31[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec11r0[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) E, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) {
	err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec11r1[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) R1 {
	r1, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec11r2[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec11r3[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2, R3 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, R3, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec11r4[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2, R3, R4 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, R3, R4, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec11r5[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2, R3, R4, R5 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec11r6[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec11r7[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec11r8[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec11r9[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec11r10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec11r11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec11r12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec11r13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec11r14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec11r15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec11r16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec11r17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec11r18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec11r19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec11r20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec11r21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec11r22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec11r23[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec11r24[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec11r25[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec11r26[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec11r27[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec11r28[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec11r29[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec11r30[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec11r31[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec12r0[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) E, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) {
	err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec12r1[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) R1 {
	r1, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec12r2[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec12r3[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2, R3 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, R3, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec12r4[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2, R3, R4 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, R3, R4, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec12r5[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2, R3, R4, R5 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec12r6[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec12r7[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec12r8[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec12r9[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec12r10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec12r11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec12r12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec12r13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec12r14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec12r15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec12r16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec12r17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec12r18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec12r19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec12r20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec12r21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec12r22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec12r23[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec12r24[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec12r25[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec12r26[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec12r27[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec12r28[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec12r29[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec12r30[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec12r31[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec13r0[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) E, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) {
	err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec13r1[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) R1 {
	r1, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec13r2[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec13r3[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2, R3 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, R3, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec13r4[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2, R3, R4 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, R3, R4, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec13r5[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2, R3, R4, R5 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec13r6[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec13r7[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec13r8[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec13r9[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec13r10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec13r11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec13r12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec13r13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec13r14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec13r15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec13r16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec13r17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec13r18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec13r19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec13r20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec13r21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec13r22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec13r23[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec13r24[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec13r25[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec13r26[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec13r27[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec13r28[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec13r29[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec13r30[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec13r31[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec14r0[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) E, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) {
	err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec14r1[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) R1 {
	r1, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec14r2[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec14r3[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2, R3 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, R3, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec14r4[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2, R3, R4 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, R3, R4, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec14r5[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2, R3, R4, R5 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec14r6[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec14r7[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec14r8[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec14r9[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec14r10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec14r11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec14r12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec14r13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec14r14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec14r15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec14r16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec14r17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec14r18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec14r19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec14r20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec14r21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec14r22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec14r23[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec14r24[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec14r25[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec14r26[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec14r27[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec14r28[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec14r29[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec14r30[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec14r31[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec15r0[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) E, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) {
	err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec15r1[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) R1 {
	r1, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec15r2[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec15r3[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2, R3 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, R3, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec15r4[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2, R3, R4 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, R3, R4, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec15r5[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2, R3, R4, R5 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec15r6[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec15r7[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec15r8[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec15r9[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec15r10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec15r11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec15r12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec15r13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec15r14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec15r15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec15r16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec15r17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec15r18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec15r19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec15r20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec15r21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec15r22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec15r23[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec15r24[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec15r25[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec15r26[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec15r27[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec15r28[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec15r29[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec15r30[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec15r31[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec16r0[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) E, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) {
	err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec16r1[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) R1 {
	r1, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec16r2[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec16r3[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2, R3 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, R3, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec16r4[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2, R3, R4 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, R3, R4, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec16r5[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2, R3, R4, R5 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec16r6[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec16r7[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec16r8[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec16r9[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec16r10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec16r11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec16r12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec16r13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec16r14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec16r15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec16r16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec16r17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec16r18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec16r19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec16r20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec16r21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec16r22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec16r23[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec16r24[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec16r25[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec16r26[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec16r27[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec16r28[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec16r29[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec16r30[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec16r31[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec17r0[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) E, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) {
	err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec17r1[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) R1 {
	r1, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec17r2[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec17r3[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2, R3 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, R3, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec17r4[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2, R3, R4 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, R3, R4, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec17r5[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2, R3, R4, R5 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec17r6[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec17r7[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec17r8[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec17r9[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec17r10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec17r11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec17r12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec17r13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec17r14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec17r15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec17r16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec17r17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec17r18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec17r19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec17r20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec17r21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec17r22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec17r23[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec17r24[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec17r25[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec17r26[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec17r27[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec17r28[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec17r29[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec17r30[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec17r31[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec18r0[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) E, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) {
	err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec18r1[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) R1 {
	r1, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec18r2[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec18r3[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2, R3 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, R3, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec18r4[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2, R3, R4 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, R3, R4, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec18r5[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2, R3, R4, R5 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec18r6[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec18r7[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec18r8[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec18r9[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec18r10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec18r11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec18r12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec18r13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec18r14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec18r15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec18r16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec18r17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec18r18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec18r19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec18r20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec18r21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec18r22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec18r23[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec18r24[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec18r25[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec18r26[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec18r27[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec18r28[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec18r29[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec18r30[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec18r31[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec19r0[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) E, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) {
	err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec19r1[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) R1 {
	r1, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec19r2[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec19r3[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2, R3 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, R3, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec19r4[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2, R3, R4 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, R3, R4, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec19r5[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2, R3, R4, R5 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec19r6[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec19r7[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec19r8[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec19r9[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec19r10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec19r11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec19r12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec19r13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec19r14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec19r15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec19r16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec19r17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec19r18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec19r19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec19r20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec19r21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec19r22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec19r23[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec19r24[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec19r25[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec19r26[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec19r27[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec19r28[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec19r29[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec19r30[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec19r31[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec20r0[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) E, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) {
	err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec20r1[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) R1 {
	r1, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec20r2[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec20r3[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2, R3 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, R3, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec20r4[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2, R3, R4 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, R3, R4, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec20r5[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2, R3, R4, R5 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec20r6[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec20r7[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec20r8[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec20r9[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec20r10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec20r11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec20r12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec20r13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec20r14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec20r15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec20r16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec20r17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec20r18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec20r19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec20r20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec20r21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec20r22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec20r23[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec20r24[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec20r25[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec20r26[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec20r27[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec20r28[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec20r29[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec20r30[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec20r31[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec21r0[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) E, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) {
	err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec21r1[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) R1 {
	r1, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec21r2[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec21r3[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2, R3 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, R3, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec21r4[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2, R3, R4 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, R3, R4, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec21r5[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2, R3, R4, R5 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec21r6[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec21r7[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec21r8[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec21r9[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec21r10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec21r11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec21r12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec21r13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec21r14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec21r15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec21r16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec21r17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec21r18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec21r19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec21r20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec21r21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec21r22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec21r23[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec21r24[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec21r25[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec21r26[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec21r27[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec21r28[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec21r29[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec21r30[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec21r31[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec22r0[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) E, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) {
	err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec22r1[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) R1 {
	r1, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec22r2[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec22r3[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2, R3 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, R3, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec22r4[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2, R3, R4 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, R3, R4, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec22r5[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2, R3, R4, R5 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec22r6[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec22r7[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec22r8[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec22r9[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec22r10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec22r11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec22r12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec22r13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec22r14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec22r15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec22r16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec22r17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec22r18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec22r19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec22r20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec22r21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec22r22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec22r23[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec22r24[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec22r25[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec22r26[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec22r27[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec22r28[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec22r29[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec22r30[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec22r31[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec23r0[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) E, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) {
	err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec23r1[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) R1 {
	r1, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec23r2[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec23r3[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2, R3 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, R3, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec23r4[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2, R3, R4 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, R3, R4, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec23r5[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2, R3, R4, R5 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec23r6[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec23r7[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec23r8[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec23r9[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec23r10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec23r11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec23r12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec23r13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec23r14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec23r15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec23r16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec23r17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec23r18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec23r19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec23r20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec23r21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec23r22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec23r23[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec23r24[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec23r25[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec23r26[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec23r27[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec23r28[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec23r29[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec23r30[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec23r31[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec24r0[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) E, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) {
	err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec24r1[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) R1 {
	r1, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec24r2[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec24r3[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2, R3 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, R3, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec24r4[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2, R3, R4 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, R3, R4, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec24r5[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2, R3, R4, R5 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec24r6[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec24r7[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec24r8[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec24r9[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec24r10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec24r11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec24r12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec24r13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec24r14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec24r15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec24r16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec24r17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec24r18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec24r19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec24r20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec24r21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec24r22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec24r23[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec24r24[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec24r25[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec24r26[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec24r27[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec24r28[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec24r29[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec24r30[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec24r31[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec25r0[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) E, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) {
	err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec25r1[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) R1 {
	r1, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec25r2[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec25r3[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2, R3 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, R3, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec25r4[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2, R3, R4 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, R3, R4, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec25r5[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2, R3, R4, R5 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec25r6[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec25r7[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec25r8[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec25r9[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec25r10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec25r11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec25r12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec25r13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec25r14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec25r15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec25r16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec25r17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec25r18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec25r19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec25r20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec25r21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec25r22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec25r23[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec25r24[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec25r25[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec25r26[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec25r27[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec25r28[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec25r29[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec25r30[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec25r31[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec26r0[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) E, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) {
	err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec26r1[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) R1 {
	r1, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec26r2[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec26r3[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2, R3 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, R3, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec26r4[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2, R3, R4 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, R3, R4, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec26r5[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2, R3, R4, R5 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec26r6[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec26r7[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec26r8[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec26r9[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec26r10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec26r11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec26r12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec26r13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec26r14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec26r15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec26r16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec26r17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec26r18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec26r19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec26r20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec26r21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec26r22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec26r23[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec26r24[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec26r25[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec26r26[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec26r27[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec26r28[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec26r29[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec26r30[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec26r31[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec27r0[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) E, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) {
	err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec27r1[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) R1 {
	r1, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec27r2[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec27r3[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2, R3 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, R3, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec27r4[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2, R3, R4 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, R3, R4, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec27r5[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2, R3, R4, R5 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec27r6[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec27r7[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec27r8[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec27r9[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec27r10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec27r11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec27r12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec27r13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec27r14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec27r15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec27r16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec27r17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec27r18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec27r19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec27r20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec27r21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec27r22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec27r23[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec27r24[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec27r25[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec27r26[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec27r27[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec27r28[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec27r29[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec27r30[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec27r31[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec28r0[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) E, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) {
	err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec28r1[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) R1 {
	r1, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec28r2[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec28r3[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2, R3 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, R3, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec28r4[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2, R3, R4 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, R3, R4, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec28r5[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2, R3, R4, R5 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec28r6[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec28r7[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec28r8[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec28r9[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec28r10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec28r11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec28r12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec28r13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec28r14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec28r15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec28r16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec28r17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec28r18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec28r19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec28r20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec28r21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec28r22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec28r23[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec28r24[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec28r25[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec28r26[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec28r27[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec28r28[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec28r29[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec28r30[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec28r31[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec29r0[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) E, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) {
	err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec29r1[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) R1 {
	r1, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec29r2[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec29r3[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2, R3 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, R3, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec29r4[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2, R3, R4 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, R3, R4, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec29r5[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2, R3, R4, R5 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec29r6[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec29r7[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec29r8[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec29r9[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec29r10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec29r11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec29r12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec29r13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec29r14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec29r15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec29r16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec29r17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec29r18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec29r19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec29r20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec29r21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec29r22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec29r23[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec29r24[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec29r25[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec29r26[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec29r27[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec29r28[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec29r29[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec29r30[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec29r31[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec30r0[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) E, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) {
	err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec30r1[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) R1 {
	r1, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec30r2[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec30r3[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2, R3 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, R3, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec30r4[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2, R3, R4 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, R3, R4, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec30r5[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2, R3, R4, R5 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec30r6[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec30r7[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec30r8[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec30r9[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec30r10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec30r11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec30r12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec30r13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec30r14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec30r15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec30r16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec30r17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec30r18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec30r19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec30r20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec30r21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec30r22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec30r23[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec30r24[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec30r25[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec30r26[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec30r27[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec30r28[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec30r29[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec30r30[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec30r31[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}

func Exec31r0[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) E, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) {
	err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return
}

func Exec31r1[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) R1 {
	r1, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1
}

func Exec31r2[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2
}

func Exec31r3[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2, R3 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, R3, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3
}

func Exec31r4[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2, R3, R4 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, R3, R4, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4
}

func Exec31r5[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2, R3, R4, R5 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, R3, R4, R5, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2, R3, R4, R5) {
	r1, r2, r3, r4, r5, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5
}

func Exec31r6[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2, R3, R4, R5, R6 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, R3, R4, R5, R6, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2, R3, R4, R5, R6) {
	r1, r2, r3, r4, r5, r6, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6
}

func Exec31r7[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2, R3, R4, R5, R6, R7 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, R3, R4, R5, R6, R7, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2, R3, R4, R5, R6, R7) {
	r1, r2, r3, r4, r5, r6, r7, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7
}

func Exec31r8[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2, R3, R4, R5, R6, R7, R8 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, R3, R4, R5, R6, R7, R8, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2, R3, R4, R5, R6, R7, R8) {
	r1, r2, r3, r4, r5, r6, r7, r8, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Exec31r9[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2, R3, R4, R5, R6, R7, R8, R9 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func Exec31r10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func Exec31r11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11
}

func Exec31r12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12
}

func Exec31r13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13
}

func Exec31r14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14
}

func Exec31r15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15
}

func Exec31r16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16
}

func Exec31r17[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17
}

func Exec31r18[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18
}

func Exec31r19[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19
}

func Exec31r20[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20
}

func Exec31r21[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21
}

func Exec31r22[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22
}

func Exec31r23[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23
}

func Exec31r24[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24
}

func Exec31r25[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25
}

func Exec31r26[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26
}

func Exec31r27[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27
}

func Exec31r28[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28
}

func Exec31r29[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29
}

func Exec31r30[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30
}

func Exec31r31[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31, R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31 any, E error](f func(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, A17, A18, A19, A20, A21, A22, A23, A24, A25, A26, A27, A28, A29, A30, A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31, E), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, a17 A17, a18 A18, a19 A19, a20 A20, a21 A21, a22 A22, a23 A23, a24 A24, a25 A25, a26 A26, a27 A27, a28 A28, a29 A29, a30 A30, a31 A31) (R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16, R17, R18, R19, R20, R21, R22, R23, R24, R25, R26, R27, R28, R29, R30, R31) {
	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31, err := f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31)
	if error(err) != nil {
		panic(Package(err))
	}
	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r30, r31
}
