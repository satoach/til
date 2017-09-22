#include <stdio.h>
#include <sys/queue.h>

#define PRINTLIST(p, h, e)                    \
    do {                                      \
        LIST_FOREACH(p, h, e)                 \
        {                                     \
            printf("%s: %d\n", p->str, p->n); \
        }                                     \
        printf("------\n");                   \
    } while (0)

struct entry { /* リストの各要素 */
    const char *str;
    LIST_ENTRY(entry) entries; /* リスト*/
    int n;
};

int main()
{
    LIST_HEAD(listhead, entry) head; /* entryを要素に持つリスト構造体 head */
    struct entry *pe;
    struct entry e0 = {.str = "e0", .n = 0};
    struct entry e1 = {.str = "e1", .n = 1};
    struct entry e2 = {.str = "e2", .n = 2};

    printf("\n"
           "##############################################################\n"
           "sys/queue.hのリスト操作のメモ\n"
           "https://linuxjm.osdn.jp/html/LDP_man-pages/man3/queue.3.htmlの\n"
           "写経 + \n"
           "##############################################################\n");

    LIST_INIT(&head);
    LIST_INSERT_HEAD(&head, &e0, entries);
    PRINTLIST(pe, &head, entries);

    LIST_INSERT_AFTER(&e0, &e2, entries);
    PRINTLIST(pe, &head, entries);

    LIST_INSERT_BEFORE(&e2, &e1, entries);
    PRINTLIST(pe, &head, entries);

    LIST_REMOVE(&e0, entries);
    PRINTLIST(pe, &head, entries);

    return 0;
}
