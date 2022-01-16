from collections import Counter


def get_most_to_least_common_element_difference(polimer: str) -> int:
    most_common_elements = Counter(polimer).most_common()
    return most_common_elements[0][1] - most_common_elements[-1][1]


def polimerize(template: str, insertions: dict, steps: int) -> str:
    polimers = list(template)
    for _ in range(steps):
        new_sequence = []
        for i in range(len(polimers) - 1):
            first, second = polimers[i], polimers[i + 1]
            new_sequence.extend([first, insertions[f"{first}{second}"]])
        new_sequence.append(polimers[-1])
        polimers = new_sequence

    return "".join(polimers)


if __name__ == '__main__':
    with open("input.txt", "r") as file:
        file_content = file.readlines()

    template = file_content[0].strip()

    pair_to_insertion = {}
    for insertion in file_content[2:]:
        match_pair, insert_polymer = insertion.strip().split(" -> ")
        pair_to_insertion[match_pair] = insert_polymer

    nth_polimer = polimerize(template, pair_to_insertion, 10)
    print(get_most_to_least_common_element_difference(nth_polimer))
