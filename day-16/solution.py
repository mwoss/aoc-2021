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


def sum_versions(payload: tuple) -> int:
    if payload[1] == 4:
        return payload[0]
    return payload[0] + sum(sum_versions(sub) for sub in payload[3])

def decode_packet(packet: str) -> tuple:
    # packet header
    version = int(packet[:3], 2)
    type_id = int(packet[3: 6], 2)
    size = 6
    sub_packets = list()

    if type_id == 4: # literal value
        for i in range(6, len(packet), 5):
            sub_packets.append(packet[i+1: i+5])
            size += 5
            if packet[i] == "0":
                break
        sub_packets = int("".join(sub_packets), 2)

    else: # operator
        size += 1
        if packet[6] == "0": # length in bits of the subpacket
            sub_size = int(packet[7: 22], 2)
            size += 15 + sub_size
            data = packet[22:]
            i = 0
            while i < sub_size:
                decoded_sub = decode_packet(data[i:])
                sub_packets.append(decoded_sub)
                i += decoded_sub[2]
        else: # packet[6] == 1, number of sub-packets
            sub_amount = int(packet[7: 18], 2)
            size += 11
            data = packet[size:]
            for i in range(sub_amount):
                decoded_sub = decode_packet(data)
                sub_packets.append(decoded_sub)
                size += decoded_sub[2]
                data = packet[size:]
    
    return version, type_id, size, sub_packets

# def decode_packet(packet: str) -> tuple:
#     version = int(packet[:3], 2)
#     type_id = int(packet[3:6], 2)
#     sub_packets = []
#     size = 6

#     if type_id == 4:
#         i = 6
#         while i < len(packet):
#             sub_packets.append(packet[i + 1:i + 5])
#             size += 5
#             if packet[i] == "0":
#                 break

#     else:
#         size += 1
#         if packet[6] == "0":
#             sub_size = int(packet[7:22], 2)
#             size += 15 + sub_size
#             data = packet[22:]
#             i = 0
#             while i < sub_size:
#                 decoded_sub = decode_packet(data[i:])
#                 sub_packets.append(decoded_sub)
#                 i += decoded_sub[2]
#         else:  # packet[6] == 1, number of sub-packets
#             sub_amount = int(packet[7: 18], 2)
#             size += 11
#             data = packet[size:]
#             for i in range(sub_amount):
#                 decoded_sub = decode_packet(data)
#                 sub_packets.append(decoded_sub)
#                 size += decoded_sub[2]
#                 data = packet[size:]

#     return version, type_id, size, sub_packets


if __name__ == '__main__':
    with open("input.txt", "r") as file:
        content = file.read().strip()

    binary_repr = ""
    for hex_val in content:
        binary_repr += hex_to_bin[hex_val]

    print(sum_versions(decode_packet(binary_repr)))
