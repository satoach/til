#include <iostream>
#include <map>
#include <vector>
#include <sstream>
#include <string.h>
#include <stdio.h>


#include "test.h"

struct flist_st {
    int id;
    char s[16];
};

int do_map(void)
{
    std::map<std::string, struct flist_st> m;

    for (int i = 0; i < 100; i++) {
        std::stringstream key;
        key << i;
        struct flist_st l = {
            .id = i,
        };
        snprintf(l.s, 15, "%03d", i);
        key << "hoge" << sizeof(i);
        m[key.str()] = l;
    }

    for (std::map<std::string, struct flist_st>::iterator it = m.begin(); it != m.end(); it++) {
        std::cout << it->first << " => " << it->second.s << std::endl;
    }

    return 0;
}

int do_split(void)
{
    std::vector<std::string> v;

    // std::string s = ",a,b,,c,";
    const char *s = ",a,b,,c,";
    std::string buf;
    std::stringstream ss;

    ss << s << ",";

    while (std::getline(ss, buf, ',')) {
        v.push_back(buf);
    }

    for (int i = 0; i < v.size(); i++) {
        printf("%d: %s\n",i ,v[i].c_str());
    }


    return 0;
}
