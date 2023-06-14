package rbtree

//node need left, right, parent, color, key, value, prev, next
type Color int

const (
	RED   Color = 1
	BLACK Color = 2
)

type Node[K, V any] struct {
	left   *Node[K, V]
	right  *Node[K, V]
	parent *Node[K, V]
	color  Color
	key    K
	value  V
}

func (n *Node[K, V]) GetKey() K {
	return n.key
}

func (n *Node[K, V]) GetValue() V {
	return n.value
}

func (n *Node[K, V]) SetValue(val V) {
	n.value = val
}

func (n *Node[K, V]) GetPrev() *Node[K, V] {
	return predecessor(n)
}

func (n *Node[K, V]) GetNext() *Node[K, V] {
	return successor(n)
}

//subtree min
func predecessor[K, V any](n *Node[K, V]) *Node[K, V] {
	if n.left != nil {
		return leftmax(n.left)
	}

	var ret *Node[K, V]

	pre := n.parent
	cur := n
	for pre != nil {
		if pre.left == cur {
			ret = pre
			break
		}
		cur = pre
		pre = pre.parent
	}
	return ret
}

//subtree max
func successor[K, V any](n *Node[K, V]) *Node[K, V] {
	if n.right != nil {
		return rightmin(n.right)
	}

	var ret *Node[K, V]

	pre := n.parent
	cur := n
	for pre != nil {
		if pre.right == cur {
			ret = pre
			break
		}
		cur = pre
		pre = pre.parent
	}
	return ret
}

func leftmax[K, V any](nLeft *Node[K, V]) *Node[K, V] {
	if nLeft.right != nil {
		nLeft = nLeft.right
	}
	return nLeft
}

func rightmin[K, V any](nRight *Node[K, V]) *Node[K, V] {
	if nRight.left != nil {
		nRight = nRight.left
	}
	return nRight
}
