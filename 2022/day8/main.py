from commonpy import fileio

from functools import reduce


def is_valid_for_given_path(value, path, input):
    for item in path:
        item_value_at_index = input[item[0]][item[1]]
        if value <= item_value_at_index:
            return False

    return True


def has_a_valid_path_or_is_edge(row: int, col: int, input) -> bool:
    if row == 0 or row == len(input) - 1:
        return True
    elif col == 0 or col == len(input[row]) - 1:
        return True

    value = input[row][col]
    return any(
        [
            is_valid_for_given_path(value, path, input)
            for path in get_paths_to_edges(row, col, input)
        ]
    )


def part1(input: list[list[int]]):
    return sum(
        [
            has_a_valid_path_or_is_edge(r, c, input)
            for r in range(len(input))
            for c in range(len(input[r]))
        ]
    )


def scenic_score_for_given_path(house_height, path, input):
    if not path:
        return 0
    trees_smaller = 0
    for cord in path:
        trees_smaller += 1
        if house_height <= input[cord[0]][cord[1]]:
            return trees_smaller

    return trees_smaller


def get_paths_to_edges(row: int, col: int, input: list[list]):
    num_rows = len(input)
    num_cols = len(input[row])
    items_up = [(r, col) for r in range(row - 1, -1, -1)]
    items_down = [(r, col) for r in range(row + 1, num_rows)]
    items_left = [(row, c) for c in range(col - 1, -1, -1)]
    items_right = [(row, c) for c in range(col + 1, num_cols)]
    return [items_up, items_down, items_left, items_right]


def get_scenic_score(row: int, col: int, input: list[list]):
    tree_height = input[row][col]
    scenic_scores_for_each_path = [
        scenic_score_for_given_path(tree_height, path, input)
        for path in get_paths_to_edges(row, col, input)
    ]
    return reduce(lambda x, y: x * y, scenic_scores_for_each_path)


def part2(input: list[str]):
    return max(
        [
            get_scenic_score(r, c, input)
            for r in range(len(input))
            for c in range(len(input[r]))
        ]
    )


if __name__ == "__main__":
    test_input = fileio.read_matrix_from_file("test_input.txt")
    input = fileio.read_matrix_from_file("input.txt")

    part_1_test_res = part1(test_input)
    print("part 1 test", part_1_test_res)

    part_1_res = part1(input)
    print("part 1 real", part_1_res)

    part_2_test_res = part2(test_input)
    print("part 2 test", part_2_test_res)

    part_2_res = part2(input)
    print("part 2 real", part_2_res)
