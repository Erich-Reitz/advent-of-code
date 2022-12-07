from commonpy import fileio
import os


def get_ls_lines(lines: list[str], index_ls_run_on: int):
    ls_output = []
    for line in lines[index_ls_run_on + 1 :]:
        if line.startswith("$"):
            break
        ls_output.append(line)

    return ls_output


def sum_directory(dirs: dict[str, str], dir: str) -> int:
    directory_items = dirs[dir]
    size = 0
    for item in directory_items:
        if item.startswith("dir"):
            item_dir_name = item.split(" ")[1]
            size += sum_directory(dirs, os.path.join(dir, item_dir_name))
        else:
            size += int(item.split(" ")[0])

    return size


def get_directory_size_info(input: list[str]):
    dirs = {}
    current_path = []
    for index, line in enumerate(input):
        if line.startswith("$"):
            command = line.split(" ")[1]
            if command == "cd":
                cd_into = line.split(" ")[2]
                if cd_into == "..":
                    current_path.pop()
                elif cd_into == "/":
                    current_path = ["/"]
                else:
                    current_path.append(cd_into)

            if command == "ls":
                ls_output = get_ls_lines(input, index)
                dirs[os.path.join(*current_path)] = ls_output
    return dirs


def part1(input: list[str]):
    dirs = get_directory_size_info(input)

    result = 0
    for dir in dirs:
        size = sum_directory(dirs, dir)
        if size <= 100000:
            result += size

    return result


def part2(input: list[str]):
    dirs = get_directory_size_info(input)

    sizes = [(sum_directory(dirs, dir), dir) for dir in dirs]
    total_space = sum_directory(dirs, "/")
    valid_deleations = [
        value for value in sizes if 70000000 - total_space + value[0] >= 30000000
    ]
    return min(valid_deleations)


if __name__ == "__main__":
    test_input = fileio.read_file_lines("test_input.txt")
    input = fileio.read_file_lines("input.txt")

    part_1_test_res = part1(test_input)
    print("part 1 test", part_1_test_res)

    part_1_res = part1(input)
    print("part 1 real", part_1_res)

    part_2_test_res = part2(test_input)
    print("part 2 test", part_2_test_res)

    part_2_res = part2(input)
    print("part 2 real", part_2_res)
