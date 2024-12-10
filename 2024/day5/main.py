def readFile(filename):
    with open(filename) as file:
        text = file.read()
        return text.split("\n\n")

def createAdjencyLst(rules: str):
    adjLst = {}
    for rule in rules.split("\n"):
        parent, child = rule.split("|")
        if parent not in adjLst:
            adjLst[parent] = set()
        adjLst[parent].add(child)
    return adjLst

def isOrderValid(pages):
    prev = pages[0]
    for page in pages[1:]:
        if page in adjLst and prev in adjLst[page]:
           return False 
        prev = page
    return True

def correctOrder(pages):
    prev = 0
    for i in range(1, len(pages)):
        if pages[i] in adjLst and pages[prev] in adjLst[pages[i]]:
             pages[i], pages[prev] = pages[prev], pages[i]
        prev = i
    return pages


if __name__ == "__main__":
    rules, updates = readFile("./2024/day5/input.txt")
    adjLst = createAdjencyLst(rules)
    res1 = 0
    res2 = 0
    for update in updates.split("\n")[:-1]:
        pages = update.split(",")
        if isOrderValid(pages):
            res1 += int(pages[len(pages) // 2])
        else:
            while not isOrderValid(pages):
                pages = correctOrder(pages)
            res2 += int(pages[len(pages) // 2])
    
    print("FIRST PART: ", res1)
    print("SECOND PART: ", res2)


