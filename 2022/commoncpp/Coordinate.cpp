#include "Coordinate.hpp"

namespace m_std {

Coordinate::Coordinate(int row, int col) {
    this->row = row;
    this->col = col;
}

Coordinate::Coordinate() {
    this->row = 0;
    this->col = 0;
}

bool Coordinate::operator==(const Coordinate &other) const {
    return this->row == other.row && this->col == other.col;
}

bool Coordinate::operator!=(const Coordinate &other) const {
    return this->row != other.row || this->col != other.col;
}

bool Coordinate::operator<(const Coordinate &other) const {
    return this->row < other.row && this->col < other.col;
}

bool Coordinate::operator>(const Coordinate &other) const {
    return this->row > other.row && this->col > other.col;
}

bool Coordinate::operator<=(const Coordinate &other) const {
    return this->row <= other.row && this->col <= other.col;
}

bool Coordinate::operator>=(const Coordinate &other) const {
    return this->row >= other.row && this->col >= other.col;
}

Coordinate Coordinate::operator+=(const Coordinate &other) {
    this->row += other.row;
    this->col += other.col;
    return *this;
}

Coordinate Coordinate::operator-=(const Coordinate &other) {
    this->row -= other.row;
    this->col -= other.col;
    return *this;
}
} // namespace m_std