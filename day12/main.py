from pythonutils.req import formatted_request

def is_valid(spring: str, cfg: str):
    w_i = 0
    count = 0
    # Add implicit separator at the end to avoid duplicate logic
    for c in spring + ".":
        if c == "#":
            count += 1
        elif c == ".":
            if count != 0:
                if w_i >= len(cfg) or count != cfg[w_i]:
                    return False
                w_i += 1
                count = 0
        # ? means the string hasn't finished so it's invalid
        else:
            return False
    return w_i == len(cfg)

def is_valid_so_far(spring: str, cfg: str):
    w_i = 0
    count = 0
    for c in spring + ".":
        if c == "#":
            count += 1
        elif c == ".":
            if count != 0:
                if w_i >= len(cfg) or count != cfg[w_i]:
                    return False
                w_i += 1
                count = 0
        # ? = we stop looking at the rest of the string, it's valid so far
        else:
            return True
    return w_i == len(cfg)

class Spring:
    def __init__(self, row: str, cfg: str):
        self.row = row
        working = cfg.split(",")
        self.working = [int(x) for x in working]
        self.cache = {}
    
    def __repr__(self):
        return f"Spring(row={self.row}, working={self.working})"

    # Slow brute force way, slow because we should stop early and not look at all permutations
    # If something isn't valid
    def possible_permutations(self):
        question_count = self.row.count("?")

        permutation_count = 0
        for i in range(2 ** question_count):
            binary = str(bin(i))[2:].zfill(question_count)
            new_spring = ""
            for c in self.row:
                if c == "?":
                    if binary[0] == "0":
                        new_spring += "."
                    else:
                        new_spring += "#"
                    binary = binary[1:]
                else:
                    new_spring += c
            if is_valid(new_spring, self.working):
                permutation_count += 1
        return permutation_count

    # Same as above but with pruning
    # Ah we probably need to shrink the string so far for the cache
    # Cache should be (string_so_far, )
    def possible_permutations_fast(self, string_so_far: str = ""):
        string_so_far = string_so_far or self.row
        if is_valid(string_so_far, self.working):
            return 1

        if not is_valid_so_far(string_so_far, self.working):
            return 0

        next_question = string_so_far.index("?")
        p1 = string_so_far[:next_question] + "." + string_so_far[next_question + 1:]
        p2 = string_so_far[:next_question] + "#" + string_so_far[next_question + 1:]
        return self.possible_permutations_fast(p1) + self.possible_permutations_fast(p2)

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
    # p_input = ["?###???????? 3,2,1"]
    print(part2(p_input))

    # print(part2(p_input))
