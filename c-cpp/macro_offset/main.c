#include <stdio.h>

struct funclist {
    void (*A)(void);
    void (*B)(void);
    void (*C)(void);
    void (*D)(void);
};

#define FUNC(name)                    \
    void func_##name(void)            \
    {                                 \
            printf("%s\n", __func__); \
    }
FUNC(A)
FUNC(B)
FUNC(C)
FUNC(D)
#undef FUNC

struct funclist func = {
    .A = &func_A, .B = NULL, .C = NULL, .D = &func_D,
};

int main(int argc, char *argv[])
{
    struct funclist *f = &func;

    printf( "\n"
            "##############################################################\n"
            "http://qiita.com/satoru_takeuchi/items/3769a644f7113f2c8040 の\n"
            "offsetをやってみる\n"
            "##############################################################\n");

#define checker(ptr, member)                    \
    ({                                          \
        if (ptr != NULL && ptr->member != NULL) \
            ptr->member();                      \
        else                                    \
            printf("NULL\n");                   \
    })

    checker(f, A);
    checker(f, B);
    checker(f, C);
    checker(f, D);
#undef checker

#define checker2(member, arg)               \
    ({                                      \
        if (f != NULL && f->member != NULL) \
            f->member(arg);                 \
        else                                \
            printf("NULL\n");               \
    })

    checker2(D, );
#undef checker2

#define offsetof(TYPE, MEMBER) ((size_t) & ((TYPE *)0)->MEMBER)
    printf("=======\n");
    printf("%lu\n", offsetof(struct funclist, A));
    printf("%lu\n", offsetof(struct funclist, B));
    printf("%lu\n", offsetof(struct funclist, C));
    printf("%lu\n", offsetof(struct funclist, D));

#define offsetofsize(TYPE, MEMBER) (offsetof(TYPE, MEMBER) + sizeof(((TYPE *)0)->MEMBER))
    printf("=======\n");
    printf("%lu\n", offsetofsize(struct funclist, A));
    printf("%lu\n", offsetofsize(struct funclist, B));
    printf("%lu\n", offsetofsize(struct funclist, C));
    printf("%lu\n", offsetofsize(struct funclist, D));
#undef offsetofsize
#undef offsetof

    return 0;
}
