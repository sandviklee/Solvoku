# Services

This layer is dedicated to the business logic of the application.

---

### Main Business Logic
This part describes all the important algorithms needed to solve a sudoku board.

#### CSP (Constraint Satisfaction Problems)
This revolves finding a solution that meets a set of constraints. 
There are three main components you need when you solve these problems:
1. **Variables:** What needs to be determined, in case of sudoku the cells that needs to be filled with numbers.
2. **Domains:** The range of potential values, in case of sudoku the integers ranging from 1-9
3. **Constraints:** The ruleset, in case of sudoku the restriction of number placement, where each row, column and box (3x3) only has one instance of each number.

We need to define three main algorithms here:

1. **Backtracking:** Depth-first search algorithm that investigates the search space of potential solutions until it finds a desired solution. It tries a variable and then backtracks if it gets stuck.
2. **Forward-Checking:** Each variable still available gets applied the restraints, and all the inconsistent values get eliminated. Basically it looks into the future, by checking each neighbor after one variable is set, and if there are no available values in the set of variables. 
3. **Propagating Constraints:** To uphold local consistency and inference to condense the search space. 

#### Sudoku Board
The sudoku board is the sudoku representation in the program.