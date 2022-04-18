def arg_list(arg_size):
    if arg_size == 0:
        return ""
    return "".join([f"a{a + 1} A{a + 1}, " for a in range(arg_size)])[:-2]


def type_list(arg_size, c):
    if arg_size == 0:
        return ""
    return "".join([f"{c}{a + 1}, " for a in range(arg_size)])[:-2]


def arg_type(arg_size):
    if arg_size == 0:
        return ""
    return type_list(arg_size, "A")


def ret_type(arg_size):
    return type_list(arg_size, "T")


for arg_size in range(32):
    AT = type_list(arg_size, "A")
    al = type_list(arg_size, "a")
    for ret_size in range(32):
        RT = type_list(ret_size, "R")
        rl = type_list(ret_size, "r")
        if ret_size > 1:
            RTC = f"({RT})"
        else:
            RTC = RT

        if arg_size == 0:
            if ret_size == 0:
                print(
                    f"func Exec{arg_size}r{ret_size}[E error](f func() (E))" " {\n "
                    f"	err := f()\n"
                    "	if error(err) != nil {\n"
                    "		panic(Package(err))\n"
                    "	}\n"
                    f"	return\n"
                    "}\n")
            else:
                print(
                    f"func Exec{arg_size}r{ret_size}[{RT} any, E error](f func() ({RT}, E)) {RTC}" " {\n "
                    f"	{rl}, err := f()\n"
                    "	if error(err) != nil {\n"
                    "		panic(Package(err))\n"
                    "	}\n"
                    f"	return {rl}\n"
                    "}\n")
        else:
            if ret_size == 0:
                print(
                    f"func Exec{arg_size}r{ret_size}[{AT} any, E error](f func({AT}) (E), {arg_list(arg_size)})" " {\n"
                    f"	err := f({al})\n"
                    "	if error(err) != nil {\n"
                    "		panic(Package(err))\n"
                    "	}\n"
                    f"	return\n"
                    "}\n")
            else:
                print(
                    f"func Exec{arg_size}r{ret_size}[{AT}, {RT} any, E error](f func({AT}) ({RT}, E), {arg_list(arg_size)}) {RTC}" " {\n "
                    f"	{rl}, err := f({al})\n"
                    "	if error(err) != nil {\n"
                    "		panic(Package(err))\n"
                    "	}\n"
                    f"	return {rl}\n"
                    "}\n")
