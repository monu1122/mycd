#include <iostream>
#include "mycd.hpp"

using std::cout;
using std::endl;



bool runTest(std::string id,MyCD cd, std::string expected)
{
    std::stringstream ss;
    std::streambuf *prevcoutbuf = std::cout.rdbuf(ss.rdbuf());
    cd.PrintDir();
    std::string text = ss.str();
    if (!text.empty() && text[text.length()-1] == '\n') {
        text.erase(text.length()-1);
    }
    int result = text.compare(expected);
    std::cout.rdbuf(prevcoutbuf);
    cout << text << endl;
    if (result == 0){
        cout << "Testcase ID: " << id << " Success" << endl;
        return true;
    }
    else {
        cout << "Testcase ID: " << id << " Failed" << endl;
        return false;
    }
}

void test()
{
    runTest("testcase_id_1",MyCD("/", "abc"), "/abc");
    runTest("testcase_id_2",MyCD("/abc/def", "ghi"), "/abc/def/ghi");
    runTest("testcase_id_3",MyCD("/abc/def", ".."), "/abc");
    runTest("testcase_id_4",MyCD("/abc/def", "/abc"), "/abc");
    runTest("testcase_id_5",MyCD("/abc/def", "/abc/klm"), "/abc/klm");
    runTest("testcase_id_6",MyCD("/abc/def", "../.."), "/");
    runTest("testcase_id_7",MyCD("/abc/def", "../../.."), "/");
    runTest("testcase_id_8",MyCD("/abc/def", "."), "/abc/def");
    runTest("testcase_id_9",MyCD("/abc/def", "..klm"), "..klm: No such file or directory");
    runTest("testcase_id_10",MyCD("/abc/def", "//////"), "/");
    runTest("testcase_id_11",MyCD("/abc/def", "......"), "......: No such file or directory");
    runTest("testcase_id_12",MyCD("/abc/def", "../gh///../klm/."), "/abc/klm");
}

int main()
{
    test();
}