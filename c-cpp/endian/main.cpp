#include <iostream>
#include <iomanip>
#include <cstdint>

using namespace std;

inline uint16_t swap_endian(uint16_t x)
{
    return ((x & 0xff) << 8) | ((x >> 8) & 0xff);
}

inline uint32_t swap_endian(uint32_t x)
{
    return (uint32_t)swap_endian((uint16_t)(x & 0xffff)) << 16 |
           (uint32_t)swap_endian((uint16_t)((x >> 16) & 0xffff));
}

inline uint64_t swap_endian(uint64_t x)
{
    return (uint64_t)swap_endian((uint32_t)(x & 0xffffffff)) << 32 |
           (uint64_t)swap_endian((uint32_t)((x >> 32) & 0xffffffff));
}

template <typename T>
void swap_data(T &x)
{
    x = swap_endian(x);
}

int main()
{
    uint16_t a = 0x0011;
    cout << hex << setw(4) << setfill('0') << a << " <-> "
         << hex << setw(4) << setfill('0') << swap_endian(a) << endl;

    uint32_t b = 0x00112233;
    cout << hex << setw(8) << setfill('0') << b << " <-> "
         << hex << setw(8) << setfill('0') << swap_endian(b) << endl;

    uint64_t c = 0x0011223344556677;
    cout << hex << setw(16) << setfill('0') << c << " <-> "
         << hex << setw(16) << setfill('0') << swap_endian(c) << endl;

    swap_data(a);
    swap_data(b);
    swap_data(c);
    cout << hex << a << " " << b << " " << c << " " << endl;

    return 0;
}
