#pragma once

namespace m_std {
    class Coordinate {
        public: 
            int row;
            int col;
            Coordinate(int row, int col);
            Coordinate();
            bool operator==(const Coordinate &other) const;
            bool operator!=(const Coordinate &other) const;
            bool operator<(const Coordinate &other) const;
            bool operator>(const Coordinate &other) const;
            bool operator<=(const Coordinate &other) const;
            bool operator>=(const Coordinate &other) const;
            Coordinate operator+=(const Coordinate &other);
            Coordinate operator-=(const Coordinate &other);
            
    };
}