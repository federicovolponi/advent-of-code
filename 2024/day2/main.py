from typing import List


def isSafe(report: List[int], i=-1) -> bool:
    if isDescreasing(report, i) or isIncreasing(report, i):
        return True
    return False


def isIncreasing(report, i=-1):
    prev = 0
    for level in report[1:]:
        if prev == i:
            prev += 1
            continue
        if report[prev] < level and 1 <= abs(report[prev] - level) <= 3:
            prev += 1
            continue
        return False
    return True


def isDescreasing(report, i=-1):
    prev = 0
    for level in report[1:]:
        if prev == i:
            prev += 1
            continue
        if report[prev] > level and 1 <= abs(report[prev] - level) <= 3:
            prev += 1
            continue
        return False
    return True


reports = []

with open("input.txt") as file:
    for line in file:
        reports.append([int(x) for x in line.rstrip().split()])

safeReports = 0
for report in reports:
    if isSafe(report):
        safeReports += 1
print("PART 1: ", safeReports)

safeReports = 0
for report in reports:
    if isSafe(report):
        safeReports += 1
    else:
        n = len(report)
        for i in range(n):
            modReport = report[:i] + report[i + 1 :]
            if isSafe(modReport):
                safeReports += 1
                break
print("PART 2: ", safeReports)
