#include "Grid.hpp"

namespace m_std {
template <typename T>
Grid<T>::Grid() {
    this->rowSize = 0;
    this->colSize = 0;
}

template <typename T>
Grid<T>::Grid(std::vector<std::vector<T>> data) {
    this->data = data;
    this->rowSize = data.size();
    this->colSize = data[0].size();
}



template <typename T>
T Grid<T>::getPoint(int row, int col) {
    return this->data[row][col];
}

template <typename T>
T Grid<T>::getPointDefault(int row, int col, T defaultValue) {
    if (row < 0 || row >= this->rowSize || col < 0 || col >= this->colSize) {
        return defaultValue;
    }
    return this->data[row][col];
}

template <typename T>
std::vector<m_std::Coordinate> Grid<T>::getNeighborsCoordinates(m_std::Coordinate coordinate) {
    std::vector<m_std::Coordinate> neighbors;
    for (int row = coordinate.row - 1; row <= coordinate.row + 1; row++) {
        for (int col = coordinate.col - 1; col <= coordinate.col + 1; col++) {
            if ((row == coordinate.row && col == coordinate.col)  || row < 0 || row >= this->rowSize || col < 0 || col >= this->colSize) {
                continue;
            }
            neighbors.push_back(m_std::Coordinate(row, col));
        }
    }
    return neighbors;
}
// template instantiations
} // namespace m_std


