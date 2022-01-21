from collections import defaultdict
from typing import Dict


def perform_insertion_process(polymer_template: str, insertions: Dict[str, str], steps: int) -> Dict[str, int]:
    element_pairs_to_count = defaultdict(lambda: 0)

    for i in range(len(polymer_template) - 1):
        element_pairs_to_count[f'{polymer_template[i]}{polymer_template[i + 1]}'] += 1

    for _ in range(steps):
        temp_element_map = defaultdict(lambda: 0)
        for element_pair, count in element_pairs_to_count.items():
            insert_element = insertions[element_pair]
            temp_element_map[f'{element_pair[0]}{insert_element}'] += count
            temp_element_map[f'{insert_element}{element_pair[1]}'] += count
        element_pairs_to_count = temp_element_map

    return element_pairs_to_count


def get_polymer_common_differance(element_pairs_to_count: Dict[str, int]) -> int:
    element_counter = defaultdict(lambda: 0)
    for element_pair, count in element_pairs_to_count.items():
        element_counter[element_pair[0]] += count
        element_counter[element_pair[1]] += count

    most_common_element = max(element_counter.values())
    least_common_element = min(element_counter.values())

    # each letter was counted twice (besides polymers with odd number of elements) as we working of pair of elements
    # which were created by moving windows of size 2, eg: NNCE -> [NN, NC, CE]
    # we have to dived that number by 2, to get expected, correct answer
    return (most_common_element - least_common_element + 1) // 2


if __name__ == '__main__':
    with open("input.txt", "r") as file:
        file_content = file.readlines()

    template = file_content[0].strip()

    pair_to_insertion = {}
    for insertion in file_content[2:]:
        match_pair, insert_polymer = insertion.strip().split(" -> ")
        pair_to_insertion[match_pair] = insert_polymer

    init_elements_count = perform_insertion_process(template, pair_to_insertion, 10)
    final_elements_count = perform_insertion_process(template, pair_to_insertion, 40)

    print(get_polymer_common_differance(init_elements_count))
    print(get_polymer_common_differance(final_elements_count))
