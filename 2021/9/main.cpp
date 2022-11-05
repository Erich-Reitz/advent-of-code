#include <vector>
#include <iostream>
#include <fstream>
#include <unordered_map>
#include <algorithm>


std::vector<std::vector<int>> readMatrixFromFile(const std::string filename) {
    std::ifstream input(filename);
    std::vector<std::vector<int>> res;
    std::string line;
    while(std::getline(input, line)) {
        std::size_t length = line.length();
        std::vector<int> row(length);
        for (size_t i = 0; i < length; i++) {
            row[i] = line[i] - '0';
        }
        res.push_back(row);
    }
    return res;
}

void markBasins(int currentRow, int currentCol, std::vector<std::vector<int>> &matrix, 
                std::vector<std::vector<bool>> &seenInBasin, 
                std::unordered_map<int, int> &basinSize, int basinNumber) {

  if (currentRow < 0 || currentRow >= matrix.size() || currentCol < 0 || currentCol >= matrix[0].size()) {
    return; 
  }
  if (matrix[currentRow][currentCol] == 9) {
    return;
  }
  if (seenInBasin[currentRow][currentCol]) {
    return;
  }
  seenInBasin[currentRow][currentCol] = true; 
  if (basinSize.find(basinNumber) != basinSize.end()) {
    basinSize[basinNumber] += 1; 
  } else {
    basinSize[basinNumber] = 1; 
  }

  markBasins(currentRow+1, currentCol, matrix, seenInBasin, basinSize, basinNumber); 
  markBasins(currentRow-1, currentCol, matrix, seenInBasin, basinSize, basinNumber); 
  markBasins(currentRow, currentCol+1, matrix, seenInBasin, basinSize, basinNumber); 
  markBasins(currentRow, currentCol-1, matrix, seenInBasin, basinSize, basinNumber); 
}

int main () {
  std::vector<std::vector<int>> matrix = readMatrixFromFile("input.txt");   
  std::vector<std::vector<bool>> seenInBasin (matrix.size(), std::vector<bool>(matrix[0].size(), 0)); 

  std::unordered_map<int, int> basinSize; 

  int seenBasins = 0 ;
  for (size_t i = 0; i < matrix.size(); i++) {
    for (size_t j = 0; j < matrix[0].size(); j++) {
      if (seenInBasin[i][j] || matrix[i][j] == 9) {
        continue;
      } else {
        markBasins(i, j, matrix, seenInBasin, basinSize, seenBasins);
        seenBasins+=1;
      }
    }
  }


  std::vector<int> basinSizesVec; 
  for (auto& it : basinSize) {
      basinSizesVec.push_back(it.second); 
  }

  std::sort(basinSizesVec.begin(), basinSizesVec.end()); 
  int ans = 1; 
  for (size_t i = basinSizesVec.size() - 3; i < basinSizesVec.size(); i++) {
    ans *= basinSizesVec[i]; 
  }

  std::cout << ans << std::endl;
}
