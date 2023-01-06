
 type Node struct {
 	val   int
 	left  *Node
 	right *Node
 }

 func (n *Node) add(val int) {

 	if val <= n.val {
 		//add left
 		if n.left != nil {
 			n.left.add(val)
 		} else {
 			n.left = &Node{
 				val: val,
 			}
 		}
 	} else {
 		// add Right
 		if n.right != nil {
 			n.right.add(val)
 		} else {
 			n.right = &Node{
 				val: val,
 			}
 		}
 	}
 }

 func (n *Node) printTree(side string) {

 	if side == "right" {
 		fmt.Print("     " + strconv.Itoa(n.val))
 	} else {
 		fmt.Print(" " + strconv.Itoa(n.val))
 	}
 	fmt.Println()
 	if n.left != nil {
 		n.left.printTree("left")
 	}
 	if n.right != nil {
 		n.right.printTree("right")
 	}
 }
