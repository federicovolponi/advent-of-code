def readFile(filename):
    with open(filename) as file:
        for line in file:
            yield line.rstrip()

def backtracking(map, curr, i, j, visited):
    if curr and map[i][j] != curr[-1] + 1:
        return
    if (i, j) not in visited and map[i][j] == 9:
        res[0] += 1
        visited.add((i, j))
        return
    curr.append(map[i][j])
    for dx, dy in [(0, 1), (1, 0), (-1, 0), (0, -1)]:
        if 0 <= i + dx < len(map) and 0 <= j + dy < len(map[0]):
            backtracking(map, curr, i + dx, j + dy, visited)
    curr.pop()

def backtrackingPart2(map, curr, i, j):
    if curr and map[i][j] != curr[-1] + 1:
        return
    if map[i][j] == 9:
        res[0] += 1
        visited.add((i, j))
        return
    curr.append(map[i][j])
    for dx, dy in [(0, 1), (1, 0), (-1, 0), (0, -1)]:
        if 0 <= i + dx < len(map) and 0 <= j + dy < len(map[0]):
            backtrackingPart2(map, curr, i + dx, j + dy)
    curr.pop()


if __name__ == "__main__":
    map = []
    for line in readFile("2024/day10/input.txt"):
        map.append([int(x) for x in line])
    res = [0]
    for i in range(len(map)):
        for j in range(len(map[0])):
            if map[i][j] == 0:
                curr = []
                visited = set()
                backtracking(map, curr, i, j, visited)
    print("FIRST PART: ", res[0])
    
    res = [0]
    for i in range(len(map)):
        for j in range(len(map[0])):
            if map[i][j] == 0:
                curr = []
                backtrackingPart2(map, curr, i, j)
    print("SECOND PART: ", res[0])
