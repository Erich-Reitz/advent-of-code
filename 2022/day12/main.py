from commonpy import fileio


def part1(input: str):
    pass


def part2(input: str):
    pass


if __name__ == "__main__":
    test_input = fileio.read_file("test_input.txt")
    input = fileio.read_file("input.txt")

    part_1_test_res = part1(test_input)
    print("part 1 test", part_1_test_res)

    part_1_res = part1(input)
    print("part 1 real", part_1_res)

    part_2_test_res = part2(test_input)
    print("part 2 test", part_2_test_res)

    part_2_res = part2(input)
    print("part 2 real", part_2_res)
