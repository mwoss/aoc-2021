from dataclasses import dataclass
from typing import List, Tuple, Set


@dataclass
class Fold:
    direction: str
    position: int


def fold_left(dots: Set[Tuple[int, int]], fold_pos: int, max_x: int) -> Set[Tuple[int, int]]:
    after_fold = set()
    for x, y in dots:
        if x >= fold_pos:
            after_fold.add((x-fold_pos, y))
        else:
            after_fold.add((abs(x-fold_pos+1), y))
    return after_fold


def fold_up(dots: Set[Tuple[int, int]], fold_pos: int, max_y: int) -> Set[Tuple[int, int]]:
    after_fold = set()
    for x, y in dots:
        if y <= fold_pos:
            after_fold.add((x, y))
        else:
            after_fold.add((abs(x - fold_pos + 1), y))
    return after_fold


def count_dots_after_fold(dots: Set[Tuple[int, int]], folds: List[Fold]) -> int:
    max_x = max(dots, key=lambda d: d[0])[0]
    max_y = max(dots, key=lambda d: d[1])[1]
    #
    # board = [[0 for _ in range(max_x)] for _ in range(max_y)]

    for fold in folds:
        if fold.direction == "x":
            dots = fold_left(dots, fold.position, max_x)
        else:
            dots = fold_up(dots, fold.position, max_y)

    return len(dots)


if __name__ == '__main__':
    dots, folds = set(), []
    with open("input.txt", "r") as file:
        for line in file:
            line = line.strip()

            if not line:
                continue
            if not line.startswith("fold"):
                x, y = line.split(",")
                dots.add((int(x), int(y)))
            else:
                direction, position = line.lstrip("fold along ").split("=")
                folds.append(Fold(direction, int(position)))

    print(count_dots_after_fold(dots, [folds[0]]))
