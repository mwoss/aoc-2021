import pprint
import sys
from typing import List, Set, Tuple

sys.setrecursionlimit(100000)


def _propagate_energy(octo_cave: List[List[int]], already_flashed: Set[Tuple[int, int]], x: int, y: int):
    neighbours = {
        (x + 1, y + 1), (x - 1, y - 1), (x + 1, y - 1), (x - 1, y + 1), (x + 1, y), (x - 1, y), (x, y + 1), (x, y - 1)
    }
    cave_size = len(octo_cave) - 1
    flashes = 0

    for neighbour in neighbours:
        xn, yn = neighbour
        if 0 > xn or xn > cave_size or 0 > yn or yn > cave_size:
            continue
        if (xn, yn) in already_flashed:
            continue

        octo_cave[xn][yn] += 1

        if octo_cave[xn][yn] > 9:
            flashes += 1
            already_flashed.add((xn, yn))
            octo_cave[xn][yn] = 0
            flashes += _propagate_energy(octo_cave, already_flashed, xn, yn)

    return flashes


def count_octopuses_flashes(octo_cave: List[List[int]], n: int) -> int:
    # simulate the flashes
    cave_size = len(octo_cave)
    flashes_count = 0
    for day in range(n):
        already_flashed = set()
        for i in range(cave_size):
            for j in range(cave_size):
                if (i, j) in already_flashed:
                    continue

                octo_cave[i][j] += 1
                if octo_cave[i][j] > 9:
                    already_flashed.add((i, j))
                    flashes_count += 1
                    octo_cave[i][j] = 0
                    flashes_count += _propagate_energy(octo_cave, already_flashed, i, j)

        if (day + 1) % 10 == 0:
            pprint.pprint(octo_cave)
            print("---------------------")

    return flashes_count

def find_first_simultaneous_flash(octo_cave: List[List[int]], threshold: int) -> int:
    # simulate the flashes
    cave_size = len(octo_cave)

    for day in range(threshold):
        already_flashed = set()
        # overall_energy = 0
        for i in range(cave_size):
            for j in range(cave_size):
                if (i, j) in already_flashed:
                    # overall_energy += octo_cave[i][j]
                    continue

                octo_cave[i][j] += 1
                if octo_cave[i][j] > 9:
                    already_flashed.add((i, j))
                    octo_cave[i][j] = 0
                    _propagate_energy(octo_cave, already_flashed, i, j)

                # overall_energy += octo_cave[i][j]

        if sum(sum(row) for row in octo_cave) == 0:
            return day + 1

    return -1


if __name__ == '__main__':
    octo_cave = []
    with open("input.txt", "r") as file:
        for octo_row in file.readlines():
            octo_cave.append([int(energy_lvl) for energy_lvl in octo_row.strip()])

    # print(count_octopuses_flashes(octo_cave, 100))
    print(find_first_simultaneous_flash(octo_cave, 1000))
