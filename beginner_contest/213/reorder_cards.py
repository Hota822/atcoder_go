from sys import stdin
import numpy as np

hwn = stdin.readline().split()

n = int(hwn[2])
num_list = [list(map(int, input().split())) for _ in range(n)]

result = np.zeros((int(hwn[0]), int(hwn[1])))
h_set = set()
w_set = set()

for idx, item in enumerate(num_list):
	result[int(item[0] - 1)][int(item[1]) - 1] = idx
	h_set.add(int(item[0]))
	w_set.add(int(item[1]))

a = np.array(result)

h_lis = [i for i in range(int(hwn[2]) - 1) if i not in h_set]
w_lis = [i for i in range(int(hwn[2]) - 1) if i not in w_set]
result = np.delete(result, h_lis, 0)
result = np.delete(result, w_lis, 1)

last = [i for i in range(int(hwn[2]) - 1)]
for i, item_h in enumerate(result):
	for j, item_w in enumerate(item_h):
		if item_w == 0:
			continue
		else
			last[item_w - 1] = i + " " + j
		
for v in last:
	print(v)
