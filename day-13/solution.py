from dataclasses import dataclass
from typing import List, Set


@dataclass(frozen=True)
class Dot:
    x: int
    y: int


@dataclass(frozen=True)
class Fold:
    direction: str
    position: int


def fold_paper(dots: Set[Dot], folds: List[Fold]) -> Set[Dot]:
    for fold in folds:
        max_x: int = max(dots, key=lambda d: d.x).x
        max_y: int = max(dots, key=lambda d: d.y).y

        if fold.direction == "x":
            dots = _fold_left(dots, fold.position, max_x)
        else:
            dots = _fold_up(dots, fold.position, max_y)

    return dots


def print_paper(dots: Set[Dot]) -> None:
    max_x = max(dots, key=lambda d: d.x).x
    max_y = max(dots, key=lambda d: d.y).y

    board = [["." for _ in range(max_x + 1)] for _ in range(max_y + 1)]

    for dot in dots:
        board[dot.y][dot.x] = "X"

    for row in board:
        print(row)


def _fold_left(dots: Set[Dot], fold_position: int, max_x: int) -> Set[Dot]:
    after_fold, len_after_fold = set(), max(max_x - fold_position, fold_position)
    for dot in dots:
        after_fold.add(Dot(len_after_fold - abs(dot.x - fold_position), dot.y))
    return after_fold


def _fold_up(dots: Set[Dot], fold_position: int, max_y: int) -> Set[Dot]:
    after_fold, len_after_fold = set(), max(max_y - fold_position, fold_position)
    for dot in dots:
        after_fold.add(Dot(dot.x, len_after_fold - abs(dot.y - fold_position)))
    return after_fold


if __name__ == '__main__':
    dots, folds = set(), []
    with open("input.txt", "r") as file:
        for line in file:
            line = line.strip()
            if not line:
                continue
            if line.startswith("fold"):
                direction, position = line.lstrip("fold along ").split("=")
                folds.append(Fold(direction, int(position)))
            else:
                x, y = line.split(",")
                dots.add(Dot(int(x), int(y)))

    single_folded_paper = fold_paper(dots, [folds[0]])
    fully_folded_paper = fold_paper(dots, folds)

    print(f"Dots after single fold: {len(single_folded_paper)}")
    print(f"Dots after all folds: {len(fully_folded_paper)}")
    print("Board after all folds:")
    print_paper(fully_folded_paper)
