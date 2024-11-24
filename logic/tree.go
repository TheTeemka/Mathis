package logic

type ExpType int64

type Node struct {
	Op    TokenType
	Value ExpType
	Left  *Node
	Right *Node
}

func newNode(op TokenType, value ExpType) *Node {
	return &Node{
		Op:    op,
		Value: value,
	}
}

func unaryNode(op TokenType, node *Node) *Node {
	if op == C_Plus {
		return node
	}
	temp := newNode(op, 0)
	temp.Left = newNode(C_Num, 0)
	temp.Right = node
	return temp
}
func compute(a, b ExpType, oper TokenType) ExpType {
	switch oper {
	case C_Plus:
		return a + b
	case C_Minus:
		return a - b
	case C_Multiply:
		return a * b
	case C_Divide:
		return a / b
	}
	return a
}

func (n *Node) Solve() ExpType {
	if n.Op == C_Num {
		return n.Value
	}
	a := n.Left.Solve()
	b := n.Right.Solve()
	return compute(a, b, n.Op)
}
