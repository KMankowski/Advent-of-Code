import re

pattern = r'^(\d+)\1+$'

with open('./input.txt', 'r') as file:
    total = 0
    for chunk in file.read().split(','):
        low = int(chunk.split('-')[0])
        high = int(chunk.split('-')[1])
        for curr_id in range(low, high + 1):
            if re.search(pattern, str(curr_id)):
                total += curr_id
    print(total)
