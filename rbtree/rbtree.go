package rbtree

// 计算空节点，将空节点视为黑色
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
type Comparator[T any] func(a, b T) int

type RbTree[K, V any] struct {
	root   *Node[K, V]
	size   int
	keyCmp Comparator[K]
}

func New[K, V any](cmp Comparator[K]) *RbTree[K, V] {
	return &RbTree[K, V]{keyCmp: cmp}
}

func (t *RbTree[K, V]) FindNode(key K) *Node[K, V] {
	check := t.root
	for check != nil {
		if t.keyCmp(key, check.GetKey()) == 0 {
			return check
		} else if t.keyCmp(key, check.GetKey()) < 0 {
			check = check.left
		} else if t.keyCmp(key, check.GetKey()) > 0 {
			check = check.right
		}
	}
	return nil
}

func (t *RbTree[K, V]) Insert(key K, value V) {
	newNode := &Node[K, V]{key: key, value: value, color: RED}
	check := t.root
	for check != nil {
		if t.keyCmp(key, check.GetKey()) == 0 {
			check.SetValue(value)
			return
		} else if t.keyCmp(key, check.GetKey()) < 0 {
			if check.left == nil {
				check.left = newNode
				newNode.parent = check
				break
			}

			check = check.left
		} else if t.keyCmp(key, check.GetKey()) > 0 {
			if check.right == nil {
				check.right = newNode
				newNode.parent = check
				break
			}
			check = check.right
		}
	}

	if check == nil {
		newNode.color = BLACK
		t.root = newNode
	} else {
		t.upInsertNode(newNode)
	}
}

func (t *RbTree[K, V]) Delete(key K) {
	check := t.FindNode(key)
	if check == nil {
		return
	}

	var del *Node[K, V]
	if check.left != nil && check.right != nil {
		del = successor(check)
	} else {
		del = check
	}

	var extra *Node[K, V]
	if del.right != nil {
		extra = del.right
	} else if del.left != nil {
		extra = del.left
	}

	if extra != nil {
		extra.parent = del.parent
	}

	extraP := del.parent
	if del.parent == nil {
		t.root = extra
	} else if del.parent.right == del {
		del.parent.right = extra
	} else {
		del.parent.left = extra
	}

	if check != del {
		check.key = del.key
		check.value = del.value
	}

	if del.color == BLACK {
		t.reBalanceDeleteNode(extra, extraP)
	}
	t.size--
}

func (t *RbTree[K, V]) reBalanceDeleteNode(extra, parent *Node[K, V]) {
	//extra不是根节点，并且extra是黑色
	for extra.parent != nil && getColor(extra) == BLACK {
		if extra != nil {
			parent = extra.parent
		}

		if parent.left == extra {
			extra = t.rbFixupLeft(parent)
		} else {

		}
	}

	if extra != nil {
		extra.color = BLACK
	}
}

func (t *RbTree[K, V]) rbFixupLeft(parent *Node[K, V]) *Node[K, V] {
	ret := t.root
	w := parent.right
	if w.color == RED {
		w.color = BLACK
		parent.color = RED
		t.leftRotate(parent)
		w = parent.left
	}

	if getColor(w.left) == BLACK && getColor(w.right) == BLACK {
		w.color = RED
		ret = parent
	} else {
		if getColor(w.right) == BLACK {
			if w.left != nil {
				w.left.color = BLACK
			}
			w.color = RED
			t.rightRotate(w)
			w = parent.right
		}
		w.color = parent.color
		parent.color = BLACK
		if w.right != nil {
			w.right.color = BLACK
		}
		t.leftRotate(parent)
	}
	return ret
}

func (t *RbTree[K, V]) upInsertNode(n *Node[K, V]) {
	cur := n
	for cur.parent != nil && cur.parent.color == RED {
		//这里隐藏的含义是 cpp 必然是存在，且必然是黑色
		curPP := cur.parent.parent
		if curPP == curPP.left {
			if curPP.right != nil && curPP.right.color == RED {
				curPP.right.color = BLACK
				curPP.color = RED
				cur.parent.color = BLACK
				cur = curPP
			} else {
				if cur == cur.parent.right {
					//这里隐藏的含义是cur仍然在左侧最下脚
					cur = cur.parent
					t.leftRotate(cur)
				}
				curPP.color = RED
				cur.parent.color = BLACK
				t.rightRotate(curPP)
			}
		} else {
			if curPP.left != nil && curPP.left.color == RED {
				curPP.left.color = BLACK
				curPP.color = RED
				cur.parent.color = BLACK
				cur = curPP
			} else {
				if cur == cur.parent.left {
					cur = cur.parent
					t.rightRotate(cur)
				}
				curPP.color = RED
				cur.parent.color = BLACK
				t.leftRotate(curPP)
			}
		}
	}
	t.root.color = BLACK
}

func (t *RbTree[K, V]) leftRotate(n *Node[K, V]) {
	left := n.left
	if left == nil {
		return
	}

	parent := n.parent
	n.parent = left
	n.left = left.right
	left.right.parent = n
	left.right = n

	if parent == nil {
		left.parent = nil
	} else {
		left.parent = parent
		if parent.right == n {
			parent.right = left
		} else {
			parent.left = left
		}
	}
}

func (t *RbTree[K, V]) rightRotate(n *Node[K, V]) {
	right := n.right
	if right == nil {
		return
	}

	parent := n.parent
	n.parent = right
	n.right = right.left
	right.left.parent = n
	right.left = n

	if parent == nil {
		right.parent = nil
	} else {
		right.parent = parent
		if parent.right == n {
			parent.right = right
		} else {
			parent.left = right
		}
	}
}

func getColor[K, V any](n *Node[K, V]) Color {
	if n == nil {
		return BLACK
	}
	return n.color
}
