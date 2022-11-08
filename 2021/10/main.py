def read_file(filename: str):
    with open(filename, "r") as f:
        return f.readlines()


def is_opening(char: str):
    return char in ["[", "(", "{", "<"]

def is_closing(char: str):
    return char in ["]", ")", "}", ">"]

def point_lookup(char: str):
    point_map = {
        ")": 3,
        "]": 57, 
        "}": 1197,
        ">": 25137
    }


    return point_map[char]


def complement(char):
    comp_map = {
        "[": "]",
        "{": "}",
        "(": ")", 
        "<": ">"
    }
    return comp_map[char]

def parse_line(line: str) -> int:
    stack = []
    for char in line:
        if is_opening(char):
            stack.append(char)

        elif is_closing(char):
            last_char = stack.pop()
            if complement(last_char) != char:
                return point_lookup(char)

    return 0


def main_pt1():
    lines  = read_file("input.txt")
    error_points = 0
    for line in lines:
        error_points_for_line = parse_line(line)
        error_points += error_points_for_line

    print(error_points)


def completion_points_for_char(char: str) -> int:
    point_map = {
        ")": 1, 
        "]": 2,
        "}": 3,
        ">": 4
    }

    return point_map[char]

def unwind_stack_for_points(stack: list):
    points = 0

    for char in stack:
        points *= 5
        points += completion_points_for_char(complement(char))
    
    return points


def get_completion_score(line: str) -> int:
    stack = []
    for char in line:
        if is_opening(char):
            stack.append(char)

        elif is_closing(char):
            last_char = stack.pop()
            if complement(last_char) != char:
                raise RuntimeError()
    
    

    return unwind_stack_for_points(reversed(stack))


def main():
    lines = read_file("input.txt")

    uncorrupted_lines = list(filter(lambda x: parse_line(x) == 0, lines))
    
    completion_points = [get_completion_score(line) for line in uncorrupted_lines]
    print(sorted(completion_points)[len(completion_points) // 2])


if __name__ == "__main__":
    main()