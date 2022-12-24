#include <iostream>
#include "mycd.hpp"

using namespace std;

int main(int argc, char *argv[])
{
    if (argc != 3)
    {
        cout << "Number of arguments does not match" << endl;
        return -1;
    }
    MyCD cd(argv[1], argv[2]);
    cd.PrintDir();

    return 0;
}