from pythonutils.req import formatted_request

def string_value(s):
    sum = 0
    for char in s:
        sum += ord(char)
        sum *= 17
        sum %= 256
    return sum

def part1(input):
    strings = input[0].split(",")
    return sum(string_value(s) for s in strings)

def part2(input):
    strings = input[0].split(",")
    final_items = {}
    # First get the final state
    for (i, s) in enumerate(strings):
        if "=" in s:
            lhs, rhs = s.split("=")
            if lhs in final_items:
                final_items[lhs][1] = int(rhs)
            else:
                final_items[lhs] = [i, int(rhs)]
        else:
            if s[:-1] in final_items:
                del final_items[s[:-1]]
    print(final_items)
    
    # Then put in boxes and do focal length calc
    boxes = [[] for _ in range(256)]
    for key, value in sorted(final_items.items(), key=lambda item: item[1][0]):
        boxes[string_value(key)].append(value[1])
    
    ret = 0
    for (j, box) in enumerate(boxes):
        for (i, lens) in enumerate(box):
            ret += (j + 1) * (i + 1) * lens
    
    return ret

if __name__ == "__main__":
    input = formatted_request(15)
    print(part1(input))
    print(part2(input))
