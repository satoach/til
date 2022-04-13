#include <stdio.h>
#include <string.h>

static unsigned int add_bit(int v)
{
    return (1 << v) - 1;
}

int main(int argc, char *argv[])
{
    for (int i = 0; i <= 32; i++) {
        printf("%02d 0x%08x\n", i, add_bit(i));
    }

    return 0;
}
