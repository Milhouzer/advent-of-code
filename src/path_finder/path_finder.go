package pathfinder

import (
	"adventofcode/src/mathematics"
	"container/heap"
	"slices"
)

type Maze struct {
	Neighbors map[Node][]Node
}

type Node struct {
	PriorityElement
	Pos Vector2
}

type Vector2 mathematics.Vector2
type Heuristic func(n Vector2) int

type AStar struct {
}

// Find the path with the lowest cost between an entry point E and an end point S.
func (a *AStar) FindPath(S, E Vector2, h Heuristic) []Vector2 {
	// Discovered nodes. Only entry node is discovered in the beginning.
	node := Node{
		PriorityElement: PriorityElement{
			Priority: 0,
		},
		Pos: E,
	}
	nodeHeap := &PriorityQueue{}
	heap.Init(nodeHeap)
	heap.Push(nodeHeap, node)

	// For node n, cameFrom[n] is the node immediately preceding it on the cheapest path from the start to n currently known.
	cameFrom := make(map[Vector2]Vector2)

	// For node n, gScore[n] is the currently known cost of the cheapest path from start to n.
	gScore := make(map[Vector2]int)

	// gScore[E] is 0 because the cost to go from the starting point to itself is 0
	gScore[E] = 0

	// For node n, fScore[n] := gScore[n] + h(n). fScore[n] represents our current best guess as to
	// how cheap a path could be from start to finish if it goes through n.
	fScore := make(map[Vector2]int)
	fScore[E] = h(E)

	for {
		if nodeHeap.Len() == 0 {
			break
		}

		current := nodeHeap.Pop().(*Node)
		if current.Pos == S {
			return reconstructPath(cameFrom, current.Pos)
		}
		/*
			// This operation can occur in O(Log(N)) time if openSet is a min-heap or a priority queue
			current := the node in openSet having the lowest fScore[] value
			if current = goal
				return reconstruct_path(cameFrom, current)

			openSet.Remove(current)
			for each neighbor of current
				// d(current,neighbor) is the weight of the edge from current to neighbor
				// tentative_gScore is the distance from start to the neighbor through current
				tentative_gScore := gScore[current] + d(current, neighbor)
				if tentative_gScore < gScore[neighbor]
					// This path to neighbor is better than any previous one. Record it!
					cameFrom[neighbor] := current
					gScore[neighbor] := tentative_gScore
					fScore[neighbor] := tentative_gScore + h(neighbor)
					if neighbor not in openSet
						openSet.add(neighbor)
		*/
	}

	return nil
}

func reconstructPath(cameFrom map[Vector2]Vector2, current Vector2) []Vector2 {
	total_path := []Vector2{current}
	for {
		current, ok := cameFrom[current]
		if !ok {
			break
		}
		total_path = append(total_path, current)
	}
	slices.Reverse(total_path)
	return total_path
}

/*
// A* finds a path from start to goal.
// h is the heuristic function. h(n) estimates the cost to reach goal from node n.
function A_Star(start, goal, h)
    // The set of discovered nodes that may need to be (re-)expanded.
    // Initially, only the start node is known.
    // This is usually implemented as a min-heap or priority queue rather than a hash-set.
    openSet := {start}

    // For node n, cameFrom[n] is the node immediately preceding it on the cheapest path from the start
    // to n currently known.
    cameFrom := an empty map

    // For node n, gScore[n] is the currently known cost of the cheapest path from start to n.
    gScore := map with default value of Infinity
    gScore[start] := 0

    // For node n, fScore[n] := gScore[n] + h(n). fScore[n] represents our current best guess as to
    // how cheap a path could be from start to finish if it goes through n.
    fScore := map with default value of Infinity
    fScore[start] := h(start)

    while openSet is not empty
        // This operation can occur in O(Log(N)) time if openSet is a min-heap or a priority queue
        current := the node in openSet having the lowest fScore[] value
        if current = goal
            return reconstruct_path(cameFrom, current)

        openSet.Remove(current)
        for each neighbor of current
            // d(current,neighbor) is the weight of the edge from current to neighbor
            // tentative_gScore is the distance from start to the neighbor through current
            tentative_gScore := gScore[current] + d(current, neighbor)
            if tentative_gScore < gScore[neighbor]
                // This path to neighbor is better than any previous one. Record it!
                cameFrom[neighbor] := current
                gScore[neighbor] := tentative_gScore
                fScore[neighbor] := tentative_gScore + h(neighbor)
                if neighbor not in openSet
                    openSet.add(neighbor)

    // Open set is empty but goal was never reached
    return failure
*/
