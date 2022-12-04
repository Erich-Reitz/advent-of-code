from commonpy import fileio

def chunks(lst, n):
    """Yield successive n-sized chunks from lst."""
    for i in range(0, len(lst), n):
        yield lst[i:i + n]

def part1(input):
    score = 0
    for line in input:
        first_half, second_half = line[0:len(line) // 2], line[len(line) // 2:]
        in_both: str = [char for char in first_half if char in second_half][0]
        if in_both.upper() == in_both:
            points = ord(in_both)-96 + 58
        else:
            points = (ord(in_both)-96)

        score += points
    
    return score

def part2(input):
    groups = list(chunks(input, 3))
    score = 0
    for group in groups:
        in_both: str = [char for char in group[0] if char in group[1] and char in group[2]][0]
        if in_both.upper() == in_both:
            points = ord(in_both)-96 + 58
        else:
            points = (ord(in_both)-96)

        score += points
    
    return score


if __name__ == "__main__":
    test_input = fileio.read_file_lines("day3_input_test.txt")
    input = fileio.read_file_lines("day3_input.txt")

    part_1_test_res = part1(test_input)
    print('part 1 test', part_1_test_res)

    part_1_res = part1(input)
    print('part 1 real', part_1_res)

    part_2_test_res = part2(test_input)
    print('part 2 test', part_2_test_res)

    part_2_res = part2(input)
    print('part 2 real', part_2_res)

