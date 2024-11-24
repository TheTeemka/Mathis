package logic

func SolveExpr(str string) int {
	parser := NewParser(str)
	root := parser.Parse()
	return int(root.Solve())
}
