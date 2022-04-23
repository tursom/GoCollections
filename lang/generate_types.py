import types_generator

for type in ["int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64"]:
    types_generator.generate(type)
