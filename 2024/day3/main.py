import re

EXPR = r"mul\(\d{1,3},\d{1,3}\)"
EXPR2 = r"mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)"
OP1 = r"\((\d+),"
OP2 = r",(\d+)\)"

def multiply(s: str)->int:
    op1 = re.findall(OP1, s)[0]
    op2 = re.findall(OP2, s)[0]
    return int(op1) * int(op2)


with open("2024/day3/input.txt") as file:
    instruction = file.read()

validMuls = re.findall(EXPR, instruction)
total = 0
for mul in validMuls:
    total += multiply(mul)
print("FIRST PART: ", total)

total = 0
muls = re.findall(EXPR2, instruction)
do = True
for instr in muls:
    if instr == "do()":
        do = True
        continue
    if instr == "don't()":
        do = False
        continue
    if do:
        total += multiply(instr)
print("SECOND PART: ", total)

