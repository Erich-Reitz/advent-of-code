from commonpy import fileio


def points_for_second_1(first, second) -> int:
    # rock
    if first == "A":
        if second == "X":
            return 1 + 3
        elif second == "Y":
            return 2 + 6
        elif second == "Z":
            return 3 + 0
    elif first == "B":
        if second == "X":
            return 1 + 0
        elif second == "Y":
            return 2 + 3
        elif second == "Z":
            return 3 + 6
    # scissors
    elif first == "C":
        if second == "X":
            return 1 + 6
        elif second == "Y":
            return 2 + 0
        elif second == "Z":
            return 3 + 3


def points_for_second_2(first, second) -> int:
    # rock
    if first == "A":
        # need to lose
        if second == "X":
            return 3 + 0
        # draw
        elif second == "Y":
            return 1 + 3
        elif second == "Z":
            return 2 + 6

    # paper
    elif first == "B":
        # lose
        if second == "X":
            return 1 + 0
        elif second == "Y":
            return 2 + 3
        elif second == "Z":
            return 3 + 6
    # scissors
    elif first == "C":
        # lose
        if second == "X":
            return 2 + 0
        # draw
        elif second == "Y":
            return 3 + 3
        elif second == "Z":
            return 1 + 6


def part1(input):
    points = 0
    for line in input:
        first, second = line.split()
        points += points_for_second_1(first, second)

    return points


def part2(input):
    points = 0
    for line in input:
        first, second = line.split()
        points += points_for_second_2(first, second)

    return points


if __name__ == "__main__":
    test_input = fileio.read_file_lines("day2_input_test.txt")
    input = fileio.read_file_lines("day2_input.txt")

    part_1_test_res = part1(test_input)
    print("part 1 test", part_1_test_res)

    part_1_res = part1(input)
    print("part 1 real", part_1_res)

    part_2_test_res = part2(test_input)
    print("part 2 test", part_2_test_res)

    part_2_res = part2(input)
    print("part 2 real", part_2_res)
