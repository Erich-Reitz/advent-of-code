from dataclasses import dataclass


@dataclass
class Octopus:
    energy_level: int
    has_flashed: bool


def get_adjacent(row: int, col: int, num_rows, num_cols, include_diagonal: bool):

    places = [
        (row + 1, col),
        (row + 1, col + 1),
        (row + 1, col - 1),
        (row, col + 1),
        (row, col - 1),
        (row - 1, col),
        (row - 1, col - 1),
        (row - 1, col + 1),
    ]

    return [
        spot
        for spot in places
        if spot[0] >= 0 and spot[0] < num_rows and spot[1] >= 0 and spot[1] < num_cols
    ]


def read_octopus_matrix():
    with open("input.txt", "r") as f:
        lines = f.readlines()

        return [[Octopus(int(char), 0) for char in line.strip()] for line in lines]


def process_flash(row, col, matrix) -> int:
    flashes_seen_here = 1
    matrix[row][col].energy_level = 0

    for (ad_row, ad_col) in get_adjacent(row, col, len(matrix), len(matrix[row]), True):
        if matrix[ad_row][ad_col].energy_level >= 9:
            flashes_seen_here += process_flash(ad_row, ad_col, matrix)
        else:
            if matrix[ad_row][ad_col].energy_level != 0:
                matrix[ad_row][ad_col].energy_level += 1

    return flashes_seen_here


def main_pt1():
    matrix = read_octopus_matrix()
    flashes = 0
    for _ in range(100):
        matrix = [
            [Octopus(energy_level=prev.energy_level + 1, has_flashed=0) for prev in row]
            for row in matrix
        ]
        for row in range(0, len(matrix)):
            for col in range(0, len(matrix[row])):
                if matrix[row][col].energy_level > 9:
                    flashes += process_flash(row, col, matrix)

    print(flashes)


def main_pt2():
    matrix = read_octopus_matrix()

    number_of_octopuses = len(matrix) * len(matrix[0])
    stage = 0
    while True:
        stage += 1
        matrix = [
            [Octopus(energy_level=prev.energy_level + 1, has_flashed=0) for prev in row]
            for row in matrix
        ]
        flashes_per_stage = 0
        for row in range(0, len(matrix)):
            for col in range(0, len(matrix[row])):
                if matrix[row][col].energy_level > 9:
                    flashes_per_stage += process_flash(row, col, matrix)
                    if flashes_per_stage == number_of_octopuses:
                        print(stage)
                        return
