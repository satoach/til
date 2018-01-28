#include <stdio.h>
#include <stdlib.h>
#include <getopt.h>

int main(int argc, char **argv)
{
    int opt;
    struct {
        int a;
        int b;
    } test = {0};
    struct option longopts[] = {
        {"t1", required_argument, NULL, '1'},
        {"t2", no_argument, &test.b, 123},
        {0, 0, 0, 0},
    };

    while ((opt = getopt_long(argc, argv, "", longopts, NULL)) != -1) {
        switch (opt) {
        case '1':
            printf("%c\n", opt);
            test.a = atoi(optarg);
            break;
        default:
            printf("no support %d\n", opt);
            break;
        }
    }

    printf("%d\n", optind);
    printf("%d %d\n", test.a, test.b);

    return 0;
}

/* vim:set sw=4 sts=4 et fenc=utf-8: */
