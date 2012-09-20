package tree

type Node struct {
    Label    int
    Children *[]Node
}

func Make(i int) *Node {
    return &Node{i, nil}
}

