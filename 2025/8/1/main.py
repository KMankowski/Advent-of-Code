import heapq
import math

coordinates = []
with open('../input.txt', 'r') as file:
    for line in file:
        coordinates.append(tuple([int(s) for s in line.strip().split(',')]))

NUMBER_OF_CONNECTIONS = 1000
connections = []
for i in range(len(coordinates)):
    for j in range(i + 1, len(coordinates)):
        first = coordinates[i]
        second = coordinates[j]
        distance = math.sqrt(((first[0] - second[0]) ** 2) + ((first[1] - second[1]) ** 2) + ((first[2] - second[2]) ** 2))
        if len(connections) < NUMBER_OF_CONNECTIONS:
            heapq.heappush(connections, (-distance, first, second))
        elif -connections[0][0] > distance:
            heapq.heappop(connections)
            heapq.heappush(connections, (-distance, first, second))

circuits = []

# reorder connections from maxheap to be smallest->largest
connections = heapq.nlargest(NUMBER_OF_CONNECTIONS, connections)

for connection in connections:
    circuit_of_first = -1
    circuit_of_second = -1
    for circuitIndex in range(len(circuits)):
        if connection[1] in circuits[circuitIndex]:
            circuit_of_first = circuitIndex
        if connection[2] in circuits[circuitIndex]:
            circuit_of_second = circuitIndex

    if circuit_of_first != -1 and circuit_of_second != -1:
        if circuit_of_first == circuit_of_second:
            continue
        circuits[circuit_of_first].update(circuits[circuit_of_second])
        del circuits[circuit_of_second]
    elif circuit_of_first == -1 and circuit_of_second == -1:
        new_circuit = set()
        new_circuit.add(connection[1])
        new_circuit.add(connection[2])
        circuits.append(new_circuit)
    elif circuit_of_first == -1:
        circuits[circuit_of_second].add(connection[1])
    elif circuit_of_second == -1:
        circuits[circuit_of_first].add(connection[2])

lengths = [len(circuit) for circuit in circuits]
largest_lengths = heapq.nlargest(3, lengths)

product = 1
for length in largest_lengths:
    product *= length

print (product)