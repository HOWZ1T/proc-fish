package fish

import rl "github.com/gen2brain/raylib-go/raylib"

// Node represents a node in a distance constrained chain.
type Node struct {
	Position           rl.Vector2
	Size               rl.Vector2
	Color              rl.Color
	DistanceConstraint float32
	Child              *Node
	Parent             *Node
}

func NewNode(position rl.Vector2, size rl.Vector2, color rl.Color, distanceConstraint float32) *Node {
	return &Node{
		Position:           position,
		Size:               size,
		Color:              color,
		DistanceConstraint: distanceConstraint,
		Child:              nil,
		Parent:             nil,
	}
}

// Draw renders the node as a circle and it's distance constraint as a green hollow circle.
func (n *Node) Draw() {
	rl.DrawCircleV(n.Position, n.Size.X, n.Color)
	rl.DrawCircleLines(int32(n.Position.X), int32(n.Position.Y), n.DistanceConstraint, rl.NewColor(0, 255, 0, 255))
}

func (n *Node) SetChild(child *Node) {
	n.Child = child
	child.Parent = n
}

func (n *Node) SetParent(parent *Node) {
	n.Parent = parent
	parent.Child = n
}

// ConstrainNodes recursively constrains the nodes in the chain based on the distance constraint of the parent.
func ConstrainNodes(parent *Node, child *Node) {
	v := rl.Vector2Subtract(child.Position, parent.Position)               // find the vector between the child and parent
	v = rl.Vector2Scale(rl.Vector2Normalize(v), parent.DistanceConstraint) // scale the vector to the distance constraint

	// find the constrained position of the child
	constrainedPos := rl.Vector2Add(parent.Position, v)
	child.Position = constrainedPos

	// draw vector at the radius of the root for debugging
	// TODO remove when no longer needed
	rl.DrawLineV(parent.Position, constrainedPos, rl.Red)

	if child.Child != nil {
		ConstrainNodes(child, child.Child)
	}
}
