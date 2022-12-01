#include <vector>
#include <string>
#include <iostream>
#include <fstream>

#include "filereads.hpp"

namespace m_std {

static std::string ERROR = "ERROR::";
static std::string INFO = "INFO::";

std::ifstream openFile(std::string filename) {
    std::ifstream input(filename);
    if (!input.is_open()) {
        std::cerr << ERROR << "Could not open " << filename << std::endl;
        exit(1);
    } else {
        std::cout << INFO << "Opened " << filename << std::endl;
    }
    return input;
}

std::vector<std::string> readLinesOfFile(std::ifstream &input) {
    /*
        * Reads all lines of a file into a vector of strings.
        * @param input: The file to read from.
        * @return: A vector of strings containing all lines of the file.
    */
    std::vector<std::string> data;
    std::string line;
    while (std::getline(input, line)) {
        data.push_back(line);
    }
    std::cout << INFO << "Read " << data.size() << " lines from file" << std::endl;
    closeFile(input);
    return data;
}

std::vector<std::string> readLinesOfFile(const std::string filename) {
    /*
        * Reads all lines of a file into a vector of strings.
        * @param input: The file to read from.
        * @return: A vector of strings containing all lines of the file.
    */
    std::ifstream input = openFile(filename);
    std::vector<std::string> data;
    std::string line;
    while (std::getline(input, line)) {
        data.push_back(line);
    }
    std::cout << INFO << "Read " << data.size() << " lines from file" << std::endl;
    closeFile(input);
    return data;
}


std::string readFile(std::ifstream &input) {
    /*
        * Reads all lines of a file into a string.
        * @param input: The file to read from.
        * @return: A string containing all lines of the file.
    */
    std::string res;
    std::string line;
    while(std::getline(input, line)) {
        res+=line;
    }
    closeFile(input);
    return res;
}

std::string readFile(const std::string filename) {
    /*
        * Reads all lines of a file into a string.
        * @param filename: The file to read from.
        * @return: A string containing all lines of the file.
    */
    std::ifstream input = openFile(filename);
    return readFile(input);
}

std::vector<std::vector<int>> readMatrixFromFile(const std::string filename) {
    /*
        * Reads a matrix from a file.
        * @param filename: The file to read from. File contains a matrix of integers
        * with no delimiters besides newlines.
        * @return: A vector of vectors containing the matrix.
    */
    std::ifstream input = openFile(filename);
    std::vector<std::vector<int>> res;
    std::string line;
    while(std::getline(input, line)) {
        std::size_t length = line.length();
        std::vector<int> row(length);
        for (int i = 0; i < length; i++) {
            row[i] = line[i] - '0';
        }
        res.push_back(row);
    }
    closeFile(input);
    return res;
}

bool failBit(std::ifstream &input) {
    /*
        * Checks if the fail bit is set.
        * @param input: The file to check.
        * @return: True if the fail bit is set, false otherwise.
    */
    return input.fail();
}

const int closeFile(std::ifstream &input) {
    /*
        * Closes a file.
        * @param input: The file to close.
        * @return: 0 if successful, 1 otherwise.
    */
    input.clear();
    input.close();
    if (failBit(input)) {
        std::cerr << ERROR << "Could not close file" << std::endl;
        return 1;
    }
    return 0;
}

} // namespace m_std
