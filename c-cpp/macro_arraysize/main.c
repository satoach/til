#include <stdio.h>


int main(int argc, char *argv[])
{
#define ARRAYSIZE(array) (sizeof(array)/sizeof(array[0]))
    int a[4];
    double b[5];
    char c[3];
    printf("%lu %lu %lu", ARRAYSIZE(a), ARRAYSIZE(b), ARRAYSIZE(c));
#undef ARRAYSIZE

    return 0;
}

