#include <stdio.h>

struct {
    int a;
    int b;
} _st = {
    .a = 1, .b = 2,
};

#define FUNC(name)                     \
    \
static int get_##name()                \
    \
{                               \
        printf("call %s\n", __func__); \
        return _st.name;               \
    \
}
FUNC(a);
FUNC(b);
#undef FUNC

int main(int argc, char *argv[])
{
    printf( "\n"
            "##############################################################\n"
            "http://qiita.com/satoru_takeuchi/items/3769a644f7113f2c8040 の\n"
            "トークン連結をやってみる\n"
            "##############################################################\n");

#define PRINT(func) printf(#func " = %d\n", func);
    PRINT(get_a());
    PRINT(get_b());
#undef PRINT

    return 0;
}
