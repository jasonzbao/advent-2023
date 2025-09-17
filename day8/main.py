from __future__ import annotations

import math
from typing import Callable
from pythonutils.req import formatted_request, format_input

example = """LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
"""

class Node:
    def __init__(self, value, left: Node, right: Node):
        self.value = value
        self.left = left
        self.right = right

def get_node_name(node: str) -> str:
    node = node.replace("(", "")
    node = node.replace(")", "")
    node = node.replace(" ", "")
    return node

def get_or_create_node(node_map: dict[str, Node], node_name: str) -> Node:
    if node_name not in node_map:
        node_map[node_name] = Node(node_name, None, None)
    return node_map[node_name]

def part1(instructions: str, node_map: dict[str, Node], start_node: str, is_end_node: Callable[[str], bool]):
    current_node = node_map[start_node]
    loops = 0
    while True:
        for (i, instruction) in enumerate(instructions):
            if instruction == "L":
                current_node = current_node.left
            elif instruction == "R":
                current_node = current_node.right
            if is_end_node(current_node.value):
                return loops * len(instructions) + i + 1
        loops += 1

def part2(instructions: str, node_map: dict[str, Node], ends_with_a: list[str]):
    steps_required = []
    for end_with_a in ends_with_a:
        steps_required.append(part1(instructions, node_map, end_with_a, lambda x: x[-1] == "Z"))
    
    # Answer will be LCM of all the steps required
    return math.lcm(*steps_required)

if __name__ == "__main__":
    input = formatted_request(8)
    # input = format_input(example)

    instructions = input[0]
    nodes = input[2:]
    
    node_map = {}
    ends_with_a = []
    for node in nodes:
        current_node = node[:3]
        lhs, rhs = node.split("=")
        parts = rhs.split(",")
        left = get_or_create_node(node_map, get_node_name(parts[0]))
        right = get_or_create_node(node_map, get_node_name(parts[1]))
        node = get_or_create_node(node_map, get_node_name(lhs))

        node.left = left
        node.right = right

        if node.value[-1] == "A":
            ends_with_a.append(node.value)

    print(part1(instructions, node_map, "AAA", lambda x: x == "ZZZ"))
    print(part2(instructions, node_map, ends_with_a))
