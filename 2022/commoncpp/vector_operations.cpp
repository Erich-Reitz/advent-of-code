#include <string>
#include <vector>

#include "vector_operations.hpp"

namespace m_std {

template <typename T>
void appendVector(std::vector<T> &v1, const std::vector<T> &v2) {
    v1.insert(v1.end(), v2.begin(), v2.end());
}

int findFirstIndexOfString(const std::vector<std::string> &lines, const std::string &str) {
    for (int i = 0; i < lines.size(); i++) {
        if (lines[i].find(str) != std::string::npos) {
            return i;
        }
    }
    return -1;
}

int findFirstIndexOfStringAfterIndex(const std::vector<std::string> &lines, const std::string &str, int index) {
    for (int i = index + 1; i < lines.size(); i++) {
        if (lines[i].find(str) != std::string::npos) {
            return i;
        }
    }
    return -1;
}
}