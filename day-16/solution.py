from dataclasses import dataclass

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

@dataclass
class PacketMetadata:
    version: int
    type_id: int
    size: int
    sub_packets: any

def sum_versions(payload: PacketMetadata) -> int:
    if payload.type_id == 4:
        return payload.version
    return payload.version + sum(sum_versions(sub) for sub in payload.sub_packets)


def decode_packet(packet: str) -> PacketMetadata:
    # packet header
    version = int(packet[:3], 2)
    type_id = int(packet[3: 6], 2)
    size = 6
    sub_packets = list()

    if type_id == 4:
        for i in range(6, len(packet), 5):
            sub_packets.append(packet[i + 1: i + 5])
            size += 5
            if packet[i] == "0":
                break
        sub_packets = int("".join(sub_packets), 2)

    else:
        size += 1
        if packet[6] == "0":
            sub_size = int(packet[7: 22], 2)
            size += 15 + sub_size
            data = packet[22:]
            i = 0
            while i < sub_size:
                decoded_sub = decode_packet(data[i:])
                sub_packets.append(decoded_sub)
                i += decoded_sub.size
        else:
            sub_amount = int(packet[7: 18], 2)
            size += 11
            data = packet[size:]
            for i in range(sub_amount):
                decoded_sub = decode_packet(data)
                sub_packets.append(decoded_sub)
                size += decoded_sub.size
                data = packet[size:]

    return PacketMetadata(version, type_id, size, sub_packets)


def payload_value(payload: PacketMetadata) -> int:
    if payload.type_id == 4:
        return payload.sub_packets

    values = [payload_value(sub) for sub in payload.sub_packets]

    if payload.type_id == 0:  # sum
        return sum(values)
    elif payload.type_id == 1:  # product
        prod = 1
        for value in values:
            prod *= value
        return prod
    elif payload.type_id == 2:  # minimum
        return min(values)
    elif payload.type_id == 3:  # maximum
        return max(values)
    elif payload.type_id == 5:  # greater than
        return 1 if values[0] > values[1] else 0
    elif payload.type_id == 6:  # less than
        return 1 if values[0] < values[1] else 0
    else:  # payload[1] == 7, equal to
        return 1 if values[0] == values[1] else 0


if __name__ == '__main__':
    with open("input.txt", "r") as file:
        content = file.read().strip()

    binary_repr = ""
    for hex_val in content:
        binary_repr += hex_to_bin[hex_val]

    print(sum_versions(decode_packet(binary_repr)))
    print(payload_value(decode_packet(binary_repr)))
