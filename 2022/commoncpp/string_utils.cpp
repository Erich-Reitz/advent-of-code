#include <string>
#include <algorithm>
#include <vector>
#include <iostream>
#include "string_utils.hpp"

namespace m_std {

void removeleadingWhitespace(std::string &s) {
    s.erase(s.begin(), std::find_if(s.begin(), s.end(), std::bind1st(std::not_equal_to<char>(), ' ')));
}

void removeLeadingChar(std::string &s, char c) {
    s.erase(s.begin(), std::find_if(s.begin(), s.end(), std::bind1st(std::not_equal_to<char>(), c)));
}

void removeTrailingWhitespace(std::string &s) {
    s.erase(std::find_if(s.begin(), s.end(), std::bind1st(std::not_equal_to<char>(), ' ')));
}

void removeBracingBrackets(std::string &str) {
    if (str.front() != '[' || str.back() != ']') {
        std::cerr << "WARNING::" << "Malformed string sent to removeBracingBrackets" << std::endl;
        throw std::runtime_error("Malformed string sent to removeBracingBrackets");
    }
    str.erase(str.end()-1);
    str.erase(str.begin());
}

std::string convertToSpaceSeperatedString(char **array, int startingIndex) {
    std::string str = "";
    for (int i = startingIndex; array[i] != nullptr; i++) {
        str += array[i];
        str += " ";
    }
    return str.substr(0, str.size() - 1);
}

std::vector<std::string> splitStringIntoLines(const std::string &str) {
    std::vector<std::string> lines;
    std::string line;
    for (char c : str) {
        if (c == '\n') {
            lines.push_back(line);
            line.clear();
        } else {
            line += c;
        }
    }
    return lines;
}

std::vector<int> stripIntsSplitByCommas(const std::string &str) {
    /*
        * Parse integers between commas in string
        * @param str: The string to parse
        * @return: A vector of ints contained by the string
    */
    std::vector<int> res;
    std::string line;
    for (char c : str) {
        if (c == ',') {
            res.push_back(atoi(line.c_str()));
            line.clear();
        } else {
            line+= c;
        }
    }
    if (line != "") {
        res.push_back(atoi(line.c_str()));
    }
    return res;
}
} // namespace m_std