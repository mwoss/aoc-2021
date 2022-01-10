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


@dataclass
class FoldedPaper:
    dots: Set[Dot]


def fold_left(dots: Set[Dot], fold_position: int, max_x: int) -> Set[Dot]:
    after_fold, len_after_fold = set(), max(max_x - fold_position, fold_position)
    for dot in dots:
        after_fold.add(Dot(len_after_fold - abs(dot.x - fold_position), dot.y))
    return after_fold


def fold_up(dots: Set[Dot], fold_position: int, max_y: int) -> Set[Dot]:
    after_fold, len_after_fold = set(), max(max_y - fold_position, fold_position)
    for dot in dots:
        after_fold.add(Dot(dot.x, len_after_fold - abs(dot.y - fold_position)))
    return after_fold


def fold_paper(dots: Set[Dot], folds: List[Fold]) -> FoldedPaper:
    for fold in folds:
        max_x: int = max(dots, key=lambda d: d.x).x
        max_y: int = max(dots, key=lambda d: d.y).y

        if fold.direction == "x":
            dots = fold_left(dots, fold.position, max_x)
        else:
            dots = fold_up(dots, fold.position, max_y)

    return FoldedPaper(dots)


def print_dotted_paper(dots: Set[Dot]) -> None:
    max_x = max(dots, key=lambda d: d.x).x
    max_y = max(dots, key=lambda d: d.y).y

    board = [["." for _ in range(max_x + 1)] for _ in range(max_y + 1)]

    for dot in dots:
        board[dot.y][dot.x] = "X"

    for row in board:
        print(row)


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

    print(fold_paper(dots, [folds[0]]))
    print(fold_paper(dots, folds))
