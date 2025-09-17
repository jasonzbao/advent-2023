from pythonutils.req import formatted_request

def get_diffs(line: list[int]) -> list[int]:
    diffs = []
    for i in range(len(line) - 1):
        diffs.append(line[i + 1] - line[i])
    return diffs

def is_zero(diffs: list[int]) -> bool:
    return all(diff == 0 for diff in diffs)

def part1(lines: list[list[int]]):
    total = 0
    for line in lines:
        stack = [line]
        while len(line) > 0:
            stack.append(get_diffs(stack[-1]))
            if is_zero(stack[-1]):
                total += sum([val[-1] for val in stack])
                break
    return total

def ret_last(stack: list[list[int]]) -> int:
    ret = 0
    for i in range(len(stack) - 1, -1, -1):
        ret = stack[i][0] - ret
    return ret

def part2(lines: list[list[int]]):
    total = 0
    for line in lines:
        stack = [line]
        while len(line) > 0:
            stack.append(get_diffs(stack[-1]))
            if is_zero(stack[-1]):
                total += ret_last(stack)
                break
    return total

if __name__ == "__main__":
    input = formatted_request(9)

    lines = [[int(x) for x in line.split(" ")] for line in input]

    print(part1(lines))
    print(part2(lines))
