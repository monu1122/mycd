#include <iostream>
#include <vector>
#include <sstream>
#include <sys/stat.h>

using namespace std;

class MyCD
{
public:
    MyCD(string cd, string nd) : currentDir(cd), newDir(nd)
    {
        formatDir();
    }

    bool PrintDir()
    {
        if (isValidDir())
        {
            cout << resultDir << endl;
            return true;
        }
        else
        {
            cout << newDir << ": No such file or directory" << endl;
            return false;
        }
    }

private:
    bool isValidDir()
    {
        struct stat st;
        return (stat(resultDir.c_str(), &st) == 0);
    }
    void formatDir()
    {
        string tempStr;
        if (newDir.c_str()[0] == '/')
        {
            tempStr = newDir;
        }
        else
        {
            tempStr = currentDir.append("/").append(newDir);
        }

        vector<string> dirList;
        std::stringstream ss(tempStr);
        string s;

        while (getline(ss, s, '/'))
        {
            if (s == "" or s == ".")
            {
                continue;
            }
            if (s == "..")
            {
                if (!dirList.empty())
                    dirList.pop_back();
                continue;
            }
            dirList.push_back(s);
        }
        if (dirList.empty())
            resultDir.append("/");
        for (auto it : dirList)
        {
            resultDir.append("/").append(it);
        }
    }

    string resultDir;
    string currentDir;
    string newDir;
};
