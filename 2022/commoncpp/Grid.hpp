#include <vector>
#include "Coordinate.hpp"

namespace m_std {
    template <typename T>
    class Grid {
        public:
            Grid();
            Grid(std::vector<std::vector<T>> data);
            T getPoint(int row, int col);
            T getPointDefault(int row, int col, T defaultValue);
            std::vector<m_std::Coordinate> getNeighborsCoordinates(m_std::Coordinate coordinate);
            int rowSize;
            int colSize;
        private:
            std::vector<std::vector<T>> data;

    };

} // namespace m_std

