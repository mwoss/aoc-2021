def polimerize(template: str, insertions: dict, steps: int) -> str:
    polimers = list(template)
    for _ in range(steps):
        new_sequence = []
        for i in range(len(polimers) - 1):
            first, second = polimers[i], polimers[i + 1]
            new_sequence.extend([first, insertions[f"{first}{second}"], second])
        polimers = new_sequence

    return "".join(polimers)


if __name__ == '__main__':
    with open("input2.txt", "r") as file:
        file_content = file.readlines()

    template = file_content[0].strip()

    pair_to_insertion = {}
    for insertion in file_content[2:]:
        match_pair, insert_polymer = insertion.strip().split(" -> ")
        pair_to_insertion[match_pair] = insert_polymer

    print(template)
    print(pair_to_insertion)

    print(polimerize(template, pair_to_insertion, 2))
