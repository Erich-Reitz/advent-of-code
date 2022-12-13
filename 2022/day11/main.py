import math
from dataclasses import dataclass
from functools import reduce

from commonpy import fileio, string_utils


@dataclass
class Monkey:
    starting_items: list[int]
    operation: str
    worry_level_divisor: int
    give_to_on_success: int
    give_to_on_failure: int
    inspected_items: int = 0

    def add_item(self, item: int) -> None:
        self.starting_items.append(item)

    def remove_item(self, item: int) -> None:
        self.starting_items.remove(item)


def parse_monkey_input(input: str) -> dict[int, Monkey]:
    blocks = input.split("\n\n")
    monkeyes = {}
    monkey_num = 0
    for block in blocks:
        for line in block.split("\n"):
            line = line.strip()
            if line.startswith("Starting items:"):
                items = string_utils.strip_integers_split_by_commas(line.split(": ")[1])
            if line.startswith("Operation:"):
                operation = line.split(": ")[1]
            if line.startswith("Test:"):
                test = int(line.split(" ").pop())
            if line.startswith("If true:"):
                given_true = int(line.split(" ").pop())
            if line.startswith("If false:"):
                given_false = int(line.split(" ").pop())
        monkeyes[monkey_num] = Monkey(items, operation, test, given_true, given_false)

        monkey_num += 1

    return monkeyes


def compute_new_worry_level(old: int, operation: str) -> int:
    evaluation_phrase = operation.split("=")[1]
    return eval(evaluation_phrase)


def part1(input: str):
    monkeys: dict[int, Monkey] = parse_monkey_input(input)

    for _ in range(0, 20):
        for monkey_num in range(0, len(monkeys)):
            current_monkey = monkeys[monkey_num]
            for item in current_monkey.starting_items[:]:
                worry_level = item
                worry_level = compute_new_worry_level(
                    worry_level, current_monkey.operation
                )
                current_monkey.inspected_items += 1
                worry_level = math.floor(worry_level / 3)
                if worry_level % current_monkey.worry_level_divisor == 0:
                    monkeys[current_monkey.give_to_on_success].add_item(worry_level)
                else:
                    monkeys[current_monkey.give_to_on_failure].add_item(worry_level)

                current_monkey.remove_item(item)

    items_inspected_per_monkeys = sorted(
        [monkey.inspected_items for monkey in monkeys.values()], reverse=True
    )
    return reduce(lambda x, y: x * y, items_inspected_per_monkeys[0:2])


def part2(input: str):
    monkeys: dict[int, Monkey] = parse_monkey_input(input)
    monkey_divisor_numbers = [monkey.worry_level_divisor for monkey in monkeys.values()]
    gcd = reduce(lambda x, y: x * y, monkey_divisor_numbers)

    for round in range(0, 10000):
        for monkey_num in range(0, len(monkeys)):
            current_monkey = monkeys[monkey_num]
            for item in current_monkey.starting_items[:]:
                worry_level = item
                worry_level = compute_new_worry_level(
                    worry_level, current_monkey.operation
                )
                current_monkey.inspected_items += 1
                if worry_level > gcd:
                    worry_level = worry_level % gcd
                if worry_level % current_monkey.worry_level_divisor == 0:
                    monkeys[current_monkey.give_to_on_success].add_item(worry_level)
                else:
                    monkeys[current_monkey.give_to_on_failure].add_item(worry_level)

                current_monkey.remove_item(item)

    items_inspected_per_monkeys = sorted(
        [monkey.inspected_items for monkey in monkeys.values()], reverse=True
    )
    return reduce(lambda x, y: x * y, items_inspected_per_monkeys[0:2])


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
