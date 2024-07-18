package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"proc-fish/fish"
)

const ScrWidth = 1280
const ScrHeight = 720

func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(ScrWidth, ScrHeight, "proc-fish")

	rl.SetTargetFPS(60)

	// create node chain
	root := fish.NewNode(rl.NewVector2(ScrWidth/2, ScrHeight/2), rl.NewVector2(10, 0), rl.Green, 40)
	nodes := []*fish.Node{root}

	prevNode := root
	var redOffset uint8 = 0
	var blueOffset uint8 = 0
	for i := 0; i < 10; i++ {
		// dynamically create nodes with different colors
		col := rl.Blue
		col.R = col.R - redOffset
		col.B = col.B - blueOffset

		blueOffset += 20
		if blueOffset > 255 {
			blueOffset = 0
			redOffset += 40
		}
		node := fish.NewNode(rl.NewVector2(ScrWidth/2, ScrHeight/2), rl.NewVector2(10, 0), col, 40)
		nodes = append(nodes, node)

		// link the nodes
		prevNode.SetChild(node)
		prevNode = node
	}

	for !rl.WindowShouldClose() {
		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			root.Position = rl.GetMousePosition()
		}

		moveVec := rl.Vector2{X: 0, Y: 0}
		if rl.IsKeyDown(rl.KeyRight) {
			moveVec.X += 1
		}

		if rl.IsKeyDown(rl.KeyLeft) {
			moveVec.X -= 1
		}

		if rl.IsKeyDown(rl.KeyUp) {
			moveVec.Y -= 1
		}

		if rl.IsKeyDown(rl.KeyDown) {
			moveVec.Y += 1
		}

		// move the root
		root.Position = rl.Vector2Add(root.Position, rl.Vector2Scale(rl.Vector2Normalize(moveVec), 5))

		rl.BeginDrawing()
		rl.ClearBackground(rl.NewColor(241, 242, 237, 255))

		fish.ConstrainNodes(root, root.Child) // TODO remove outside of drawing section when the debug render is removed

		// render the nodes
		for _, node := range nodes {
			node.Draw()
		}
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
