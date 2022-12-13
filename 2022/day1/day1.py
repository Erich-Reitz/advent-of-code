from commonpy import fileio, string_utils, list_utils


def part1(input: list[str]):
    divided_list = list_utils.split_list_by_value(input, "")
    return max(
        [
            sum(string_utils.convert_list_of_strings_to_list_ints(cal_list))
            for cal_list in divided_list
        ]
    )


def part2(input: list[str]):
    divided_list = list_utils.split_list_by_value(input, "")
    cal_totals = [
        sum(string_utils.convert_list_of_strings_to_list_ints(cal_list))
        for cal_list in divided_list
    ]
    cal_totals = sorted(cal_totals, reverse=True)
    return sum(cal_totals[0:3])


if __name__ == "__main__":
    input = fileio.read_file_lines("day1_input.txt")

    part1_result = part1(input)
    print(part1_result)
    part2_result = part2(input)
    print(part2_result)
