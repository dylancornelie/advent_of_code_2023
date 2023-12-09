from functools import reduce
import re

file = open("./input.txt", "r")
inputs = file.read().split("\n")

# inputs = [
#     "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
#     "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
#     "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
#     "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
#     "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
# ]


def solve_part_1(input: str) -> int:
    predicat = {"R": 12, "G": 13, "B": 14}
    left, right = input.split(":")
    max_cube_count = {"R": 0, "G": 0, "B": 0}
    set_of_cube_revealed = right.split(";")
    for cube_revealed in set_of_cube_revealed:
        cubes_showed = [
            match[0] for match in re.findall("(\d+ (blue|red|green))", cube_revealed)
        ]
        # Compute the max value of each cube before Elf put back cube into the bag
        local_cube_count = {"R": 0, "G": 0, "B": 0}
        for cube_showed in cubes_showed:
            if "red" in cube_showed:
                count = int(re.findall("\d+", cube_showed)[0])
                if local_cube_count["R"] < count:
                    local_cube_count["R"] = local_cube_count["R"] + count
            if "green" in cube_showed:
                count = int(re.findall("\d+", cube_showed)[0])
                if local_cube_count["G"] < count:
                    local_cube_count["G"] = local_cube_count["G"] + count
            if "blue" in cube_showed:
                count = int(re.findall("\d+", cube_showed)[0])
                if local_cube_count["B"] < count:
                    local_cube_count["B"] = local_cube_count["B"] + count
        # Update the max value
        for key, value in local_cube_count.items():
            if max_cube_count[key] < value:
                max_cube_count[key] = value

    for key, value in predicat.items():
        if max_cube_count[key] > value:
            return 0

    game_id = int(re.findall("\d+", left)[0])
    return game_id


def solve_part_2(input: str) -> int:
    _, right = input.split(":")
    max_cube_count = {"R": 0, "G": 0, "B": 0}
    set_of_cube_revealed = right.split(";")
    for cube_revealed in set_of_cube_revealed:
        cubes_showed = [
            match[0] for match in re.findall("(\d+ (blue|red|green))", cube_revealed)
        ]
        # Compute the max value of each cube before Elf put back cube into the bag
        local_cube_count = {"R": 0, "G": 0, "B": 0}
        for cube_showed in cubes_showed:
            if "red" in cube_showed:
                count = int(re.findall("\d+", cube_showed)[0])
                if local_cube_count["R"] < count:
                    local_cube_count["R"] = local_cube_count["R"] + count
            if "green" in cube_showed:
                count = int(re.findall("\d+", cube_showed)[0])
                if local_cube_count["G"] < count:
                    local_cube_count["G"] = local_cube_count["G"] + count
            if "blue" in cube_showed:
                count = int(re.findall("\d+", cube_showed)[0])
                if local_cube_count["B"] < count:
                    local_cube_count["B"] = local_cube_count["B"] + count
        # Update the min value
        for key, value in local_cube_count.items():
            if max_cube_count[key] < value:
                max_cube_count[key] = value

    power = 1
    for value in max_cube_count.values():
        power *= value

    return power


result = reduce(lambda acc, result: acc + solve_part_2(result), inputs, 0)
print(result)
