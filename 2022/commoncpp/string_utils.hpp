#pragma once

namespace m_std {
    void removeleadingWhitespace(std::string &s);

    void removeLeadingChar(std::string &s, char c);

    std::string convertToSpaceSeperatedString(char **array, int startingIndex);

    std::vector<std::string> splitStringIntoLines(const std::string &str);

    std::vector<int> stripIntsSplitByCommas(const std::string &str);

    void removeBracingBrackets(std::string &str);
}