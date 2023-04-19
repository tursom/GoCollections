#pragma once

#include<stdint.h>

uint32_t hash(uint8_t* data, size_t size, uint32_t seed) {
    uint32_t hash = seed;
    switch (size % sizeof(uint32_t)) {
    case 3:
        hash = hash*31 ^ *data++;
    case 2:
        hash = hash*31 ^ *data++;
    case 1:
        hash = hash*31 ^ *data++;
    }

    int n = size / sizeof(uint32_t);
    for (int i = 0; i < n; i++) {
        hash = hash*31 ^ *(uint32_t*)data;
        data += sizeof(uint32_t);
    }

    return hash;
}
