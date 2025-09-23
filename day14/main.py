from pythonutils.req import formatted_request

example = """O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#...."""

class Board:
    def __init__(self, board: list[str]):
        self.board = [[x for x in line] for line in board]

        self.directions = {
            'north': (-1, 0),
            'south': (1, 0),
            'west': (0, -1),
            'east': (0, 1)
        }

        self.iteration_orders = {
            'north': {'vary': 0, 'vary_range': range(len(self.board)), 'vary_step': 1,
                     'fixed': 1, 'fixed_range': range(len(self.board[0])), 'fixed_step': 1},
            'south': {'vary': 0, 'vary_range': range(len(self.board)-1, -1, -1), 'vary_step': -1,
                     'fixed': 1, 'fixed_range': range(len(self.board[0])), 'fixed_step': 1},
            'west': {'vary': 1, 'vary_range': range(len(self.board[0])), 'vary_step': 1,
                    'fixed': 0, 'fixed_range': range(len(self.board)), 'fixed_step': 1},
            'east': {'vary': 1, 'vary_range': range(len(self.board[0])-1, -1, -1), 'vary_step': -1,
                    'fixed': 0, 'fixed_range': range(len(self.board)), 'fixed_step': 1}
        }

    def move_rock(self, old_pos: tuple[int, int], new_pos: tuple[int, int]):
        old_item = self.board[old_pos[0]][old_pos[1]]
        self.board[old_pos[0]][old_pos[1]] = self.board[new_pos[0]][new_pos[1]]
        self.board[new_pos[0]][new_pos[1]] = old_item

    def find_most_rock(self, pos: tuple[int, int], direction: str):
        """Find the most distant available position in the given direction."""
        dr, dc = self.directions[direction]
        r, c = pos

        while 0 <= r < len(self.board) and 0 <= c < len(self.board[0]):
            if self.board[r][c] != ".":
                # Return the position just before this obstacle
                return (r - dr, c - dc)
            r += dr
            c += dc

        # Hit boundary, return the boundary position
        r -= dr
        c -= dc
        return (r, c)

    def slide(self, direction: str):
        """Slide all rocks in the given direction."""
        order = self.iteration_orders[direction]

        for vary_val in order[ 'vary_range']:
            for fixed_val in order['fixed_range']:
                if order['vary'] == 0:  # varying row
                    i, j = vary_val, fixed_val
                else:  # varying column
                    i, j = fixed_val, vary_val

                if self.board[i][j] == "O":
                    # Find the starting position for search (one step in the slide direction)
                    dr, dc = self.directions[direction]
                    start_pos = (i + dr, j + dc)
                    target_pos = self.find_most_rock(start_pos, direction)
                    self.move_rock((i, j), target_pos)

    # Convenience methods
    def slide_north(self):
        self.slide('north')

    def slide_south(self):
        self.slide('south')

    def slide_west(self):
        self.slide('west')

    def slide_east(self):
        self.slide('east')

    def __repr__(self):
        return "\n".join(["".join(row) for row in self.board])

    def get_weight(self):
        ret = 0
        for i in range(len(self.board)):
            for j in range(len(self.board[i])):
                if self.board[i][j] == "O":
                    ret += len(self.board) - i
        return ret

if __name__ == "__main__":
    board = Board(formatted_request(14))
    # board = Board(example.split("\n"))
    # print(board)
    # board.slide_north()
    # print(board)
    # Part 1
    # print(board.get_weight())

    # part 2
    rotations = 1000000000
    directions = ["north", "west", "south", "east"]
    cycle_length = None
    cycle_last = None
    cache = {}
    for i in range(rotations):
        cache_key = board.__repr__()
        if cache_key in cache:
            cycle_length = i - cache[cache_key][0]
            cycle_last = i
            break
        else:
            cache[cache_key] = (i, board.get_weight())
        
        for direction in directions:
            board.slide(direction)
    
    # Now that we know the cycle length, we can figure out the weight
    cycle_weights = [0 for _ in range(cycle_length)]
    for key, value in cache.items():
        first_element = (cycle_last - cycle_length)
        if value[0] >= first_element:
            cycle_weights[value[0] - first_element] = value[1]

    print(cycle_weights[(rotations - cycle_last) % cycle_length])
