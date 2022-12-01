#pragma once
#include <vector> 
#include <string>
#include <fstream>

namespace m_std {

    std::vector<std::string> readLinesOfFile(std::ifstream &input);
    std::vector<std::string> readLinesOfFile(const std::string filename); 

    std::string readFile(std::ifstream &input);

    std::string readFile(const std::string filename);

    int findFirstIndexOfString(const std::vector<std::string> &lines, const std::string &str);

    int findFirstIndexOfStringAfterIndex(const std::vector<std::string> &lines, const std::string &str, int index);

    std::vector<std::vector<int>> readMatrixFromFile(std::string filename);

    bool failBit(std::ifstream &input);
    const int closeFile(std::ifstream &input);

}