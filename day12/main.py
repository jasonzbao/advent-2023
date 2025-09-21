from pythonutils.req import formatted_request

class Spring:
    def __init__(self, row: str, cfg: str):
        self.row = row
        working = cfg.split(",")
        self.working = [int(x) for x in working]
        self.cache = {}
    
    def __repr__(self):
        return f"Spring(row={self.row}, working={self.working})"

    def cache(func):
        def wrapper(self, *args, **kwargs):
            key = args
            if key in self.cache:
                return self.cache[key]
            result = func(self, *args, **kwargs)
            self.cache[key] = result
            return result
        return wrapper
    
    @cache
    def possible_permutations_fastest(self, i: int = 0, j : int = 0, count_so_far: int = 0):
        if i >= len(self.row):
            if j >= len(self.working):
                return 1
            if j == len(self.working) - 1:
                return 1 if count_so_far == self.working[j] else 0
            return 0

        if j >= len(self.working):
            return 1 if "#" not in self.row[i:] else 0
                
        if self.row[i] == "#":
            return self.possible_permutations_fastest(i + 1, j, count_so_far + 1)
        if self.row[i] == ".":
            if count_so_far == 0:
                return self.possible_permutations_fastest(i + 1, j, count_so_far)
            else:
                if count_so_far != self.working[j]:
                    return 0
                return self.possible_permutations_fastest(i + 1, j + 1, 0)
        
        if self.row[i] == "?":
            result = 0
            # If we hit the end, we have to use a .
            if count_so_far == self.working[j]:
                result += self.possible_permutations_fastest(i + 1, j + 1, 0)
            else:
                # If it's not equal and it's 0, we could use a .
                if count_so_far == 0:
                    result += self.possible_permutations_fastest(i + 1, j, count_so_far)
                # In all cases, we could use a #
                result += self.possible_permutations_fastest(i + 1, j, count_so_far + 1)
            return result

def part1(p_input):
    springs = []
    for line in p_input:
        row, cfg = line.split(" ")
        springs.append(Spring(row, cfg))
    return sum(spring.possible_permutations_fastest() for spring in springs)

def part2(p_input):
    springs = []
    for line in p_input:
        row, cfg = line.split(" ")
        s_row, s_cfg = row, cfg
        for _ in range(4):
            s_row += "?" + row
            s_cfg += "," + cfg
        springs.append(Spring(s_row, s_cfg))
    return sum(spring.possible_permutations_fastest() for spring in springs)

if __name__ == "__main__":
    p_input = formatted_request(12)
    # print(part1(p_input))
    print(part2(p_input))
