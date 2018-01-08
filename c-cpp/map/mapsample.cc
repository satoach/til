#include <iostream>
#include <map>

template <typename T> T add(T a, T b)
{
    return a + b;
}

template <> float add(float a, float b)
{
    return a - b;
}

static int do_map_auto()
{
    std::map<std::string, int> m1;

    m1["3"] = 300;
    m1["1"] = 100;
    m1["2"] = 200;

    for (auto &x : m1) {
        std::cout << x.first << " => " << x.second << std::endl;
    }

    return 0;
}

static int do_map()
{
    std::map<std::string, int> m1;

    m1["a"] = 300;
    m1["z"] = 100;
    m1["d"] = 200;

    for (std::map<std::string, int>::iterator it = m1.begin(); it != m1.end(); it++) {
        std::cout << it->first << " => " << it->second << std::endl;
    }

    return 0;
}

int main(int argc, char const *argv[])
{
    do_map_auto();
    do_map();
    std::cout << "add(1, 2) " << add(1, 2) << std::endl;
    float a = 1.0, b = 2.0;
    std::cout << "add(1.0, 2.0) " << add(a, b) << std::endl;

    return 0;
}
