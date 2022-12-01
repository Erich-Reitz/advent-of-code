#include <iostream>
#include <string> 
#include "filereads.hpp"

int main() {
    // open hello.txt
    std::string contents = m_std::readFile("hello.txt");
    std::cout << contents << std::endl;

    return 0;
}