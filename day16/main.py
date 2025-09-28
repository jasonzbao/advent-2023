from pythonutils.req import formatted_request, format_input

directions = {
    "up": (-1, 0),
    "down": (1, 0),
    "left": (0, -1),
    "right": (0, 1)
}

dummy_input = r""".|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|...."""

class Puzzle:
    def __init__(self, input):
        self.input = input

    def print_grid(self, ret_cache):
        for i in range(len(self.input)):
            for j in range(len(self.input[0])):
                if (i, j) in ret_cache:
                    print("#", end="")
                else:
                    print(".", end="")
            print()

    def travel(self, x, y, direction):
        # we've already been here in this direction, just return
        cache = set()
        stack = [(x, y, direction)]
        while len(stack) > 0:
            x, y, direction = stack.pop()
            if (x, y, direction) in cache:
                continue

            if x < 0 or x >= len(self.input) or y < 0 or y >= len(self.input[0]):
                continue
            cache.add((x, y, direction))

            dx, dy = directions[direction]
            if self.input[x][y] == ".":
                stack.append((x + dx, y + dy, direction))
                continue
            elif self.input[x][y] == "|":
                if direction == "left" or direction == "right":
                    stack.append((x+1, y, "down"))
                    stack.append((x-1, y, "up"))
                elif direction == "up" or direction == "down":
                    stack.append((x + dx, y + dy, direction))
                continue
            elif self.input[x][y] == "-":
                if direction == "up" or direction == "down":
                    stack.append((x, y - 1 , "left"))
                    stack.append((x, y + 1, "right"))
                elif direction == "left" or direction == "right":
                    stack.append((x + dx, y + dy, direction))
                continue
            elif self.input[x][y] == "/":
                if direction == "up":
                    stack.append((x, y + 1, "right"))
                elif direction == "down":
                    stack.append((x, y - 1, "left"))
                elif direction == "left":
                    stack.append((x + 1, y, "down"))
                elif direction == "right":
                    stack.append((x - 1, y, "up"))

            elif self.input[x][y] == "\\":
                if direction == "up":
                    stack.append((x, y - 1, "left"))
                elif direction == "down":
                    stack.append((x, y + 1, "right"))
                elif direction == "left":
                    stack.append((x - 1, y, "up"))
                elif direction == "right":
                    stack.append((x + 1, y, "down"))
                continue
            else:
                print("Unknown character: ", self.input[x][y])
                return
        
        ret_cache = set()
        for (x, y, direction) in cache:
            ret_cache.add((x, y))
        # self.print_grid(ret_cache)
        return len(ret_cache), cache
      

if __name__ == "__main__":
    p_input = formatted_request(16)
    # p_input = format_input(dummy_input)

    p_input = [[x for x in line] for line in p_input]
    puzzle = Puzzle(p_input)
    print(puzzle.travel(0, 0, "right")[0])

    # part 2
    entry_points = []
    for i in range(len(p_input)):
        entry_points.append((i, 0, "right"))
        entry_points.append((i, len(p_input[0]) - 1, "left"))
    for j in range(len(p_input[0])):
        entry_points.append((0, j, "down"))
        entry_points.append((len(p_input) - 1, j, "up"))

    purge = set()
    max_ret = 0
    # We can purge points where we have visited the same point in the same direction
    for (x, y, direction) in entry_points:
        if (x, y, direction) in purge:
            continue

        ret, seen = puzzle.travel(x, y, direction)
        max_ret = max(max_ret, ret)

        for (x, y, direction) in seen:
            purge.add((x, y, direction))

    print(max_ret)