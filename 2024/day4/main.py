from itertools import product

DIRECTIONS = list(product([0, 1, -1], repeat=2))[1:]

def readFile(filename):
    with open(filename) as file:
        for line in file:
            yield line.rstrip()


def dfs(i, j, direction = None):
    if (
        (len(curr) == 2 and curr[-1] != "M")
        or (len(curr) == 3 and curr[-1] != "A")
        or (len(curr) == 4 and curr[-1] != "S")
        or (len(curr) > 4)
    ):
        return
    if len(curr) == 4:
        res[0] += 1
        return
    if direction:
        if 0 <= i + direction[0] < len(wordsGrid) and 0 <= j + direction[1] < len(wordsGrid[0]):
            curr.append(wordsGrid[i+direction[0]][j+direction[1]])
            dfs(i+direction[0], j+direction[1], direction)
            curr.pop()
        return

    for dx, dy in DIRECTIONS:
        if 0 <= i + dx < len(wordsGrid) and 0 <= j + dy < len(wordsGrid[0]):
            curr.append(wordsGrid[i+dx][j+dy])
            dfs(i+dx, j+dy, (dx, dy))
            curr.pop()

def isXMAS(i, j):
    if wordsGrid[i - 1][j - 1] == "M" and wordsGrid[i + 1][j + 1] == "S":
        if (wordsGrid[i - 1][j + 1] == "M" and wordsGrid[i + 1][j - 1] == "S") or (wordsGrid[i - 1][j + 1] == "S" and wordsGrid[i + 1][j - 1] == "M"):
            return True
    if wordsGrid[i - 1][j - 1] == "S" and wordsGrid[i + 1][j + 1] == "M":
        if (wordsGrid[i - 1][j + 1] == "M" and wordsGrid[i + 1][j - 1] == "S") or (wordsGrid[i - 1][j + 1] == "S" and wordsGrid[i + 1][j - 1] == "M"):
            return True
    return False



if __name__ == "__main__":
    wordsGrid = []
    for line in readFile("./2024/day4/input.txt"):
        wordsGrid.append([letter for letter in line])
    
    res = [0]
    for i in range(len(wordsGrid)):
        for j in range(len(wordsGrid[0])):
            if wordsGrid[i][j] == "X":
                curr = ["X"]
                dfs(i, j)
    print("FIRST PART: ", res[0])
    
    res = [0]
    for i in range(1, len(wordsGrid) - 1):
        for j in range(1, len(wordsGrid[0]) - 1):
            if wordsGrid[i][j] == "A" and isXMAS(i, j):
                res[0] += 1
    print("SECOND PART: ", res[0])
