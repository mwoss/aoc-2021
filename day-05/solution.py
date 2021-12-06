from collections import defaultdict


def get_amount_hydrothermal_vent_clusters() -> int:
    with open("input.txt", "r") as file:
        raw_lines = file.read()

    diagram = defaultdict(lambda: 0)

    for raw_line in raw_lines.strip().split("\n"):
        cords = raw_line.split(" -> ")
        x1, y1 = cords[0].split(",")
        x2, y2 = cords[1].split(",")
        x1, x2, y1, y2 = int(x1), int(x2), int(y1), int(y2)

        if x1 == x2:
            for yi in range(min(y1, y2), max(y1, y2) + 1):
                diagram[f"{x1},{yi}"] += 1
        elif y1 == y2:
            for xi in range(min(x1, x2), max(x1, x2) + 1):
                diagram[f"{xi},{y1}"] += 1

    return sum(1 for amount in diagram.values() if amount >= 2)


def get_amount_hydrothermal_vent_clusters_diagonal() -> int:
    with open("input.txt", "r") as file:
        raw_lines = file.read()

    diagram = defaultdict(lambda: 0)

    for raw_line in raw_lines.strip().split("\n"):
        cords = raw_line.split(" -> ")
        x1, y1 = cords[0].split(",")
        x2, y2 = cords[1].split(",")
        x1, x2, y1, y2 = int(x1), int(x2), int(y1), int(y2)

        if x1 == x2:
            for yi in range(min(y1, y2), max(y1, y2) + 1):
                diagram[f"{x1},{yi}"] += 1
        elif y1 == y2:
            for xi in range(min(x1, x2), max(x1, x2) + 1):
                diagram[f"{xi},{y1}"] += 1
        else:
            x_sign = 1 if x2 - x1 > 0 else -1
            y_sign = 1 if y2 - y1 > 0 else -1
            for offset in range(abs(x2 - x1) + 1):
                diagram[f"{x1 + x_sign * offset},{y1 + y_sign * offset}"] += 1

    return sum(1 for amount in diagram.values() if amount >= 2)


if __name__ == '__main__':
    print(get_amount_hydrothermal_vent_clusters())
    print(get_amount_hydrothermal_vent_clusters_diagonal())
