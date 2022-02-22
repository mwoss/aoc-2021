hex_to_bin = {
    "0": "0000",
    "1": "0001",
    "2": "0010",
    "3": "0011",
    "4": "0100",
    "5": "0101",
    "6": "0110",
    "7": "0111",
    "8": "1000",
    "9": "1001",
    "A": "1010",
    "B": "1011",
    "C": "1100",
    "D": "1101",
    "E": "1110",
    "F": "1111",
}

if __name__ == '__main__':
    with open("input.txt", "r") as file:
        content = file.read().strip()

    binary_repr = ""
    for hex_val in content:
        binary_repr += hex_to_bin[hex_val]

    print(binary_repr)
