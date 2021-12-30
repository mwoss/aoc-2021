from copy import deepcopy
from typing import List, Set, Tuple

MIN_ADMISSIBLE_ENERGY = 0
MAX_ADMISSIBLE_ENERGY = 9


def count_octopuses_flashes(octo_cave: List[List[int]], n: int) -> int:
    flashes_count, cave_size = 0, len(octo_cave)

    for _ in range(n):
        already_flashed = set()
        for x in range(cave_size):
            for y in range(cave_size):
                if (x, y) in already_flashed:
                    continue

                flashes_count += _increment_octo_energy(x, y, octo_cave, already_flashed)

    return flashes_count


def find_first_simultaneous_flash(octo_cave: List[List[int]], threshold: int) -> int:
    cave_size = len(octo_cave)

    for day in range(threshold):
        already_flashed = set()
        for x in range(cave_size):
            for y in range(cave_size):
                if (x, y) in already_flashed:
                    continue

                _increment_octo_energy(x, y, octo_cave, already_flashed)

        if sum(sum(row) for row in octo_cave) == 0:
            return day + 1

    return -1


def _increment_octo_energy(x: int, y: int, octo_cave: List[List[int]], already_flashed: Set[Tuple[int, int]]):
    if (x, y) in already_flashed:
        return 0

    octo_cave[x][y] += 1

    if octo_cave[x][y] > MAX_ADMISSIBLE_ENERGY:
        octo_cave[x][y] = MIN_ADMISSIBLE_ENERGY
        already_flashed.add((x, y))
        return _propagate_energy(x, y, octo_cave, already_flashed) + 1

    return 0


def _propagate_energy(x: int, y: int, octo_cave: List[List[int]], already_flashed: Set[Tuple[int, int]]):
    flashes_count, cave_size = 0, len(octo_cave) - 1
    neighbours = [
        (x + 1, y + 1), (x - 1, y - 1), (x + 1, y - 1),
        (x - 1, y + 1), (x + 1, y), (x - 1, y), (x, y + 1), (x, y - 1)
    ]

    for neighbour in neighbours:
        xn, yn = neighbour

        if 0 > xn or xn > cave_size or 0 > yn or yn > cave_size:
            continue

        flashes_count += _increment_octo_energy(xn, yn, octo_cave, already_flashed)

    return flashes_count


if __name__ == '__main__':
    octo_cave = []
    with open("input.txt", "r") as file:
        for octo_row in file.readlines():
            octo_cave.append([int(energy_lvl) for energy_lvl in octo_row.strip()])

    # we have to use deepcopy as we mutate our input ://
    print(count_octopuses_flashes(deepcopy(octo_cave), 100))
    print(find_first_simultaneous_flash(deepcopy(octo_cave), 1000))
