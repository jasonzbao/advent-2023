import itertools

from pythonutils.req import formatted_request, format_input

example = """...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#....."""

scaling_factor = 999999

def print_grid(grid):
    for row in grid:
        print("".join(row))

def empty_columns_and_rows(grid):
    empty_row = []
    empty_column = []
    for i in range(len(grid)):
        if all(grid[i][j] == "." for j in range(len(grid[i]))):
            empty_row.append(i)
        if all(grid[j][i] == "." for j in range(len(grid[i]))):
            empty_column.append(i)
    return empty_row, empty_column

def expand_empties(grid, empty_row, empty_column):
    print(empty_row, empty_column)
    for i in empty_row[::-1]:
        grid = grid[:i] + [["." for _ in range(len(grid[0]))]] + grid[i:]
    for i in empty_column[::-1]:
        for j in range(len(grid)):
            grid[j] = grid[j][:i] + ["."] + grid[j][i:]
    return grid

def distance(a, b):
    return abs(a[0] - b[0]) + abs(a[1] - b[1])

def part1(grid):
    galaxies = get_galaxy_coords(grid)

    ret = 0
    for combo in itertools.combinations(galaxies, 2):
        ret += distance(combo[0], combo[1])
    return ret

def get_galaxy_coords(grid):
    galaxies = []
    for i in range(len(grid)):
        for j in range(len(grid[i])):
            if grid[i][j] == "#":
                galaxies.append([i, j])
    return galaxies

def get_new_coords(coord, empty_rows, empty_columns):
    x = coord[0] + scaling_factor * len([i for i in empty_rows if i < coord[0]])
    y = coord[1] + scaling_factor * len([i for i in empty_columns if i < coord[1]])
    return [x, y]


def part2(grid, empty_row, empty_column):
    galaxies = get_galaxy_coords(grid)
    ret = 0
    for combo in itertools.combinations(galaxies, 2):
        ret += distance(
            get_new_coords(combo[0], empty_row, empty_column),
            get_new_coords(combo[1], empty_row, empty_column))
    return ret

if __name__ == "__main__":
    grid = formatted_request(11)
    # grid = format_input(example)

    og_grid = [[x for x in line] for line in grid]

    empty_row, empty_column = empty_columns_and_rows(og_grid)
    grid = expand_empties(og_grid, empty_row, empty_column)
    print(part1(grid))

    # expanding the arrays is probably not the best way to do this, instead i'll find the coords of the galaxies
    # and do math to find their new distance
    print(part2(og_grid, empty_row, empty_column))
