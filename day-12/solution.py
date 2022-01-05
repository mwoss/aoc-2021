from collections import defaultdict
from typing import Dict, List, Set, Optional, Callable

START_NODE_ID = "start"
END_NODE_ID = "end"


def count_unique_paths(cave: Dict[str, List[str]], path_counter: Callable) -> int:
    if START_NODE_ID not in cave or END_NODE_ID not in cave:
        raise Exception("One of require nodes (start/end) are missing in given cave representation")
    return path_counter(START_NODE_ID, cave, {START_NODE_ID})


def _count_unique_paths_single_entrance(node: str, cave: Dict[str, List[str]], path: Set[str]) -> int:
    if node == END_NODE_ID:
        return 1

    count = 0
    for next_node in cave[node]:
        if next_node.isupper() or next_node not in path:
            count += _count_unique_paths_single_entrance(next_node, cave, {*path, next_node})

    return count


def _count_unique_paths_twice_entrance(
    node: str,
    cave: Dict[str, List[str]],
    path: Set[str],
    visited_twice: Optional[str] = None
) -> int:
    if node == END_NODE_ID:
        return 1

    count = 0
    for next_node in cave[node]:
        if next_node.isupper() or next_node not in path:
            count += _count_unique_paths_twice_entrance(next_node, cave, {*path, next_node}, visited_twice)
        elif next_node != START_NODE_ID and visited_twice is None:
            count += _count_unique_paths_twice_entrance(next_node, cave, {*path, next_node}, next_node)

    return count


if __name__ == '__main__':
    cave = defaultdict(list)
    with open("input.txt") as file:
        for line in file:
            start_node, end_node = line.strip().split("-")
            cave[start_node].append(end_node)
            cave[end_node].append(start_node)

    print(count_unique_paths(cave, _count_unique_paths_single_entrance))
    print(count_unique_paths(cave, _count_unique_paths_twice_entrance))
