from collections import Counter, defaultdict


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


def polimerize_cache(template: str, insertions: dict, steps: int) -> int:
    element_map = defaultdict(lambda: 0)
    letter_map = defaultdict(lambda: 0)
    for i in range(len(template) - 1):
        element_map[f'{template[i]}{template[i + 1]}'] += 1

    for _ in range(steps):
        temp_element_map = defaultdict(lambda: 0)
        temp_letter_map = defaultdict(lambda: 0)
        for element, count in element_map.items():
            letter = insertions[element]
            temp_element_map[f'{element[0]}{letter}'] += count
            temp_element_map[f'{letter}{element[1]}'] += count
            temp_letter_map[element[0]] += count
            temp_letter_map[element[1]] += count
            temp_letter_map[letter] += count

        element_map = temp_element_map
        letter_map = temp_letter_map

    most_common_element = max(letter_map.values())
    least_common_element = min(letter_map.values())
    return most_common_element - least_common_element


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

    print(polimerize_cache(template, pair_to_insertion, 10))
