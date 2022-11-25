#  Copyright (c) 2022 tursom. All rights reserved.
#  Use of this source code is governed by a GPL-3
#  license that can be found in the LICENSE file.

def generate(base_type: str):
    if base_type[0] == 'u':
        i = 2
    else:
        i = 1
    package_type = base_type[0:i].upper() + base_type[i:]

    lines = []
    # package
    lines.append("package lang\n")

    # imports
    lines.append("import \"strconv\"\n")

    # type define
    lines.append(f"type {package_type} {base_type}\n")

    if base_type == "int32":
        lines.append("type Rune = Int32\n")

    # As(Type) interface
    lines.append(f"type As{package_type} interface ""{\n"
                 "	Object\n"
                 f"	As{package_type}() {package_type}\n"
                 "}\n")

    # Cast(Type) function
    lines.append(f"func Cast{package_type}(v any) ({base_type}, bool) ""{\n"
                 "	switch i := v.(type) {\n"
                 f"	case {base_type}:\n"
                 "		return i, true\n"
                 f"	case As{package_type}:\n"
                 f"		return i.As{package_type}().V(), true\n"
                 "	default:\n"
                 "		return 0, false\n"
                 "	}\n"
                 "}\n")

    # Equals(Type) function
    lines.append(f"func Equals{package_type}(i1 As{package_type}, i2 any) bool ""{\n"
                 f"	i2, ok := Cast{package_type}(i2)\n"
                 f"	return ok && i2 == i1.As{package_type}().V()\n"
                 "}\n")

    # type.V() function
    lines.append(f"func (i {package_type}) V() {base_type} ""{\n"
                 f"	return {base_type}(i)\n"
                 "}\n")

    # type.P() function
    lines.append(f"func (i *{package_type}) P() *{base_type} ""{\n"
                 f"	return (*{base_type})(i)\n"
                 "}\n")

    # As(Type) interface implement
    lines.append(f"func (i {package_type}) As{package_type}() {package_type} ""{\n"
                 f"	return i\n"
                 "}\n")

    # String
    if base_type[0] == 'u':
        lines.append(f"func (i {package_type}) String() string ""{\n"
                     f"	return strconv.FormatUint(uint64(i), 10)\n"
                     "}\n")
    else:
        lines.append(f"func (i {package_type}) String() string ""{\n"
                     f"	return strconv.FormatInt(int64(i), 10)\n"
                     "}\n")

    # AsObject
    lines.append(f"func (i {package_type}) AsObject() Object ""{\n"
                 f"	return i\n"
                 "}\n")

    # Equals
    lines.append(f"func (i {package_type}) Equals(e Object) bool ""{\n"
                 f"	return Equals{package_type}(i, e)\n"
                 "}\n")

    # ToString
    lines.append(f"func (i {package_type}) ToString() String ""{\n"
                 "	return NewString(i.String())\n"
                 "}\n")

    # HashCode
    if package_type.endswith("64"):
        lines.append(f"func (i {package_type}) HashCode() int32 ""{\n"
                     "	return Hash64(&i)\n"
                     "}\n")
    else:
        lines.append(f"func (i {package_type}) HashCode() int32 ""{\n"
                     "	return int32(i)\n"
                     "}\n")

    # Compare
    lines.append(f"func (i {package_type}) Compare(t {package_type}) int ""{\n"
                 "	switch {\n"
                 "	case i > t:\n"
                 "		return 1\n"
                 "	case i == t:\n"
                 "		return 0\n"
                 "	default:\n"
                 "		return -1\n"
                 "	}\n"
                 "}\n")

    with open(f"{package_type}.go", "w", encoding="utf-8") as out:
        out.writelines([f"{l}\n" for l in lines])
