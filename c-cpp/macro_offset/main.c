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

#define dataof(TYPE, MEMBER) (((TYPE *)0)->MEMBER)
#define membersizeof(TYPE, MEMBER) sizeof(dataof(TYPE, MEMBER))
#define offsetof(TYPE, MEMBER) ((size_t)&dataof(TYPE, MEMBER))
    printf("=======\n");
    printf("%lu\n", offsetof(struct funclist, A));
    printf("%lu\n", offsetof(struct funclist, B));
    printf("%lu\n", offsetof(struct funclist, C));
    printf("%lu\n", offsetof(struct funclist, D));

#define offsetofsize(TYPE, MEMBER) (offsetof(TYPE, MEMBER) + membersizeof(TYPE, MEMBER))
    printf("=======\n");
    printf("%lu\n", offsetofsize(struct funclist, A));
    printf("%lu\n", offsetofsize(struct funclist, B));
    printf("%lu\n", offsetofsize(struct funclist, C));
    printf("%lu\n", offsetofsize(struct funclist, D));

#include <string.h>
    {
        struct Prop {
            char a;
            int b;
            short c;
            char d[7];
        };
        struct Prop org = {
            .a = '1',
            .b = 2,
            .c = 3,
            .d = "d = d",
        };
        struct Prop prop = {0};
        printf("=======\n");
        memcpy(&prop.a, (void*)&org + offsetof(struct Prop, a), membersizeof(struct Prop, a));
        memcpy(&prop.b, (void*)&org + offsetof(struct Prop, b), membersizeof(struct Prop, b));
        memcpy(&prop.c, (void*)&org + offsetof(struct Prop, c), membersizeof(struct Prop, c));
        memcpy(&prop.d, (void*)&org + offsetof(struct Prop, d), membersizeof(struct Prop, d));

        printf("%lu \t %lu \t %c\n", offsetof(struct Prop, a), membersizeof(struct Prop, a), prop.a);
        printf("%lu \t %lu \t %d\n", offsetof(struct Prop, b), membersizeof(struct Prop, b), prop.b);
        printf("%lu \t %lu \t %d\n", offsetof(struct Prop, c), membersizeof(struct Prop, c), prop.c);
        printf("%lu \t %lu \t %s\n", offsetof(struct Prop, d), membersizeof(struct Prop, d), prop.d);
    }

#undef offsetofsize
#undef offsetof

    return 0;
}
