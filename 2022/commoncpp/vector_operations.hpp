# pragma once

namespace m_std {
    template <typename T>
    void appendVector(std::vector<T> &v1, const std::vector<T> &v2);

    int findFirstIndexOfString(const std::vector<std::string> &lines, const std::string &str);

    int findFirstIndexOfStringAfterIndex(const std::vector<std::string> &lines, const std::string &str, int index);

}