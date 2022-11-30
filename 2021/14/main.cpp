#include <algorithm>
#include <iostream>
#include <string>
#include <map>

#include "commoncpp/filereads.hpp"
#include "gmpxx.h"

using n_pair = std::pair<char, char>  ;
using n_pair_count_map = std::map<std::pair<char, char>, mpz_class>  ;


template<typename KeyType, typename ValueType>
std::pair<KeyType,ValueType> get_max( const std::map<KeyType,ValueType>& x ) {
    using pairtype=std::pair<KeyType,ValueType>;
    return *std::max_element(x.begin(), x.end(), [] (const pairtype & p1, const pairtype & p2) {
        return p1.second < p2.second;
    });
}

template<typename KeyType, typename ValueType>
std::pair<KeyType,ValueType> get_min( const std::map<KeyType,ValueType>& x ) {
    using pairtype=std::pair<KeyType,ValueType>;
    return *std::max_element(x.begin(), x.end(), [] (const pairtype & p1, const pairtype & p2) {
        return p1.second > p2.second;
    });
}

n_pair_count_map pairsOfPolymers(std::string strand) {
    n_pair_count_map result;
    for (size_t i = 0; i < strand.size() - 1; i++) {
        result[std::make_pair(strand[i], strand[i+1])] ++;
    }
    return result;
}

n_pair_count_map processPairInsertionRules(n_pair_count_map  splitPairs, std::map<n_pair, char> insertionRules) {
    n_pair_count_map resultMap = {};
    for (auto splitPairMapEntry : splitPairs) {
        n_pair splitPair = splitPairMapEntry.first;
        auto occurances = splitPairMapEntry.second;
        if (insertionRules.find(splitPair) != insertionRules.end()) {
            n_pair leftPair = std::make_pair(splitPair.first, insertionRules[splitPair]);
            n_pair rightPair = std::make_pair(insertionRules[splitPair], splitPair.second );
            resultMap[leftPair] += occurances;
            resultMap[rightPair] += occurances;
        }  else {
            resultMap[splitPair] += occurances;
        }
    }
    return resultMap;
}

std::map<n_pair, char> convertRulesToMap(std::vector<std::string> rules) {
    std::map<n_pair, char> ruleMap = {};
    for (std::string rule : rules) {
        ruleMap[std::make_pair(rule[0], rule[1])] =  rule[6] ;
    }
    return ruleMap;
}

int main() {
    auto lines = m_std::readLinesOfFile("input.txt");
    auto polymerTemplate = lines[0];
    auto pairInsertionRules = std::vector<std::string>(lines.begin() + 2, lines.end());
    auto pairInsertionRuleMap = convertRulesToMap(pairInsertionRules);
    auto splitPairs = pairsOfPolymers(polymerTemplate);
    for (int step = 0; step < 40; step++) {
        splitPairs = processPairInsertionRules(splitPairs, pairInsertionRuleMap);
    }
    char lastLetter = polymerTemplate.back();
    std::map<char, mpz_class> freqMap;
    for (auto entry : splitPairs) {
        freqMap[entry.first.first] += entry.second;
    }
    freqMap[lastLetter] += 1;
    std::cout << get_max(freqMap).second - get_min(freqMap).second << std::endl;
}