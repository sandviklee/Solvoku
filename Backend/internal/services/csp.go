package services

import "math"

type CSP struct {
	variables   []variable
	domains     map[variable][]int
	constraints map[variable][]variable
}

func (ctx *CSP) InitCSPWithSudoku(sudokuCtx sudokuCtx) {
	ctx.variables = sudokuCtx.variables
	ctx.domains = sudokuCtx.domains
	ctx.constraints = sudokuCtx.constraints
}

func Solve(ctx CSP) (solution map[variable]int) {
	assignment := make(map[variable]int)
	solution = backtrack(assignment, ctx)

	return
}

func backtrack(assignment map[variable]int, ctx CSP) (result map[variable]int) {
	if len(assignment) == len(ctx.variables) {
		return assignment
	}

	v := selectUnassignedVariable(assignment, ctx)
	for _, value := range getDomainValues(v, ctx) {
		if isConsistent(v, assignment, value, ctx) {
			assignment[v] = value
			result = backtrack(assignment, ctx)
			if result != nil {
				return result
			}
			delete(assignment, v)
		}
	}
	return nil

}

// Returns the unnasiagned variable with the smallest domain.
func selectUnassignedVariable(assignment map[variable]int, ctx CSP) (v variable) {
	unassignedVariables := make([]variable, 0)

	for _, ctxVar := range ctx.variables {
		inArr := false
		for assVar, _ := range assignment {
			if ctxVar == assVar {
				inArr = true
			}
		}
		if !inArr {
			unassignedVariables = append(unassignedVariables, ctxVar)
		}
	}

	currSmall := math.MaxInt
	for _, unassVar := range unassignedVariables {
		domLen := len(ctx.domains[unassVar])
		if domLen >= currSmall {
			continue
		}
		currSmall = domLen
		v = unassVar
	}

	return
}

func getDomainValues(v variable, ctx CSP) []int {
	return ctx.domains[v]
}

func isConsistent(v variable, assignment map[variable]int, value int, ctx CSP) bool {
	for _, constraintVar := range ctx.constraints[v] {
		inArr := false
		for assignmentVar, _ := range assignment {
			if constraintVar == assignmentVar {
				inArr = true
			}
		}
		if inArr && assignment[constraintVar] == value {
			return false
		}
	}
	return true
}
