package graph

import "math"

type service struct {
	repository Repository
}

type Service interface {
	GetGraph(input *InputPoint) ([]Point, string)
}

func InitGraphService(repository Repository) *service {
	return &service{repository: repository}
}

type Point struct {
	X, Y int
}

type Node struct {
	Point
	G, H, F float64
	Parent  *Node
}

func (s *service) GetGraph(input *InputPoint) ([]Point, string) {
	start := Point{input.StartX, input.EndX}
	goal := Point{input.StartY, input.EndY}

	obstacles := []Point{
		{1, 2},
		{2, 2},
		{3, 2},
		{3, 3},
		{3, 4},
		{4, 1},
		{5, 2},
	}

	path := findPath(start, goal, obstacles)

	return path, ""
}

func findPath(start, goal Point, obstacles []Point) []Point {
	openSet := make(map[Point]*Node)
	closedSet := make(map[Point]*Node)

	startNode := &Node{Point: start}
	startNode.G = 0
	startNode.H = heuristic(start, goal)
	startNode.F = startNode.G + startNode.H

	openSet[start] = startNode

	for len(openSet) > 0 {
		var current *Node
		for _, node := range openSet {
			if current == nil || node.F < current.F {
				current = node
			}
		}

		if current.Point == goal {
			return reconstructPath(current)
		}

		delete(openSet, current.Point)
		closedSet[current.Point] = current

		for _, neighbor := range getNeighbors(current.Point, obstacles) {
			if _, exists := closedSet[neighbor]; exists {
				continue
			}

			g := current.G + distance(current.Point, neighbor)
			h := heuristic(neighbor, goal)
			f := g + h

			if _, exists := openSet[neighbor]; !exists || g < openSet[neighbor].G {
				neighborNode := &Node{
					Point:  neighbor,
					G:      g,
					H:      h,
					F:      f,
					Parent: current,
				}
				openSet[neighbor] = neighborNode
			}
		}
	}

	return nil
}

func reconstructPath(current *Node) []Point {
	var path []Point
	for current != nil {
		path = append(path, current.Point)
		current = current.Parent
	}
	// Reverse the path
	for i := 0; i < len(path)/2; i++ {
		j := len(path) - i - 1
		path[i], path[j] = path[j], path[i]
	}
	return path
}

func getNeighbors(point Point, obstacles []Point) []Point {
	// Implement logic to get neighboring points based on your obstacle constraints
	// For simplicity, let's assume 4-connected grid
	neighbors := []Point{
		{point.X + 1, point.Y},
		{point.X - 1, point.Y},
		{point.X, point.Y + 1},
		{point.X, point.Y - 1},
	}

	// Remove obstacles from neighbors
	for i := 0; i < len(neighbors); {
		neighbor := neighbors[i]
		if contains(obstacles, neighbor) {
			neighbors = append(neighbors[:i], neighbors[i+1:]...)
		} else {
			i++
		}
	}
	return neighbors
}

func contains(points []Point, point Point) bool {
	for _, p := range points {
		if p == point {
			return true
		}
	}
	return false
}

func heuristic(a, b Point) float64 {
	return distance(a, b)
}

func distance(a, b Point) float64 {
	return math.Sqrt(float64((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y)))
}
