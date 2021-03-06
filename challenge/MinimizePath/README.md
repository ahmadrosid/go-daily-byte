# Minimize Path
This question is asked by Google.

Given an NxM matrix, grid, where each cell in the matrix represents the cost of stepping on the current cell, return the minimum cost to traverse from the top-left hand corner of the matrix to the bottom-right hand corner.

Note: You may only move down or right while traversing the grid.

Ex: Given the following grid

```bash
grid = [
    [1,1,3],
    [2,3,1],
    [4,6,1]
], return 7.
The path that minimizes our cost is 1->1->3->1->1 which sums to 7.
```