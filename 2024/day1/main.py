from typing import Iterator


def readFile(filename: str) -> Iterator[str]:
    with open(filename) as file:
        for line in file:
            yield line.rstrip()


if __name__ == "__main__":
    lst1, lst2 = [], []
    for line in readFile("./input.txt"):
        value1, value2 = line.split()
        lst1.append(int(value1))
        lst2.append(int(value2))

    lst1.sort()
    lst2.sort()

    res = 0
    for val1, val2 in zip(lst1, lst2):
        res += abs(val2 - val1)
    print("PART 1: ", res)

    hashMap = {val: 0 for val in lst2}
    for val2 in lst2:
        hashMap[val2] += 1
    res = 0
    for val1 in lst1:
        res += val1 * hashMap.get(val1, 0)
    print("PART 2: ", res)
