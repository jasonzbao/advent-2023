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
    
    def move_rock(self, old_pos: tuple[int, int], new_pos: tuple[int, int]):
        old_item = self.board[old_pos[0]][old_pos[1]]
        self.board[old_pos[0]][old_pos[1]] = self.board[new_pos[0]][new_pos[1]]
        self.board[new_pos[0]][new_pos[1]] = old_item

    def find_north_most_rock(self, pos: tuple[int, int]):
        for i in range(pos[0], -1, -1):
            if self.board[i][pos[1]] != ".":
                return (i + 1, pos[1])
        return (0, pos[1])

    def slide_north(self):
        for i in range(len(self.board)):
            for j in range(len(self.board[i])):
                if self.board[i][j] == "O":
                    north_most_rock = self.find_north_most_rock((i - 1, j))
                    self.move_rock((i, j), north_most_rock)

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
    board.slide_north()
    # print(board)
    print(board.get_weight())
