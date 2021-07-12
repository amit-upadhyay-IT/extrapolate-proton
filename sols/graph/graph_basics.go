package graph

import (
	"strconv"
	"strings"
)

/* NOTE BEGIN
 * WWW is a good example of graph. Social media can also be modeled as graphs, where people can be represented as node,
 * and their connections can be represented as edges in the graph.
 * So, basically a graph is a represented with set of verticals and ends.

 * A basica difference between tree and graph is a tree is always connected, but a graph does not necessary need to be
 * connected. For example we may have a graph like this:

 *     O----------O                 O-------O
 *     |          |                / \       \
 *     |          |               /   \       \
 *     O----------O              O----O-------O

 * Let me use a terminology here:
 * Basically, the process of visiting each node in a given graph, I will call it as traversal
 * And process of vising possible nodes from a node in graph, I will call it as search.

 * Search technique can be used to implement traversal, how? coz from each nodes in the input we will do search, and
 * eventually we will end of vising all the nodes in the graph.

 * The most important algorithms on graph are all about traversal and search.

 Basic overview of breadth first search:
 * First we search for all the elements which are at a distance 1 from the current node, then we search for elements
 * which are at a shortest distance of 2 from current node and so on. This is kind of same as level order traversal
 * (not exactly same) in a tree.

 Basic overview of depth first search:
 * In DFS you go as deep as possible. From one node you go to another connected node then to another connected node, etc
 * Both DFS and BFS are equally important when it comes to problem solving.

 * Before discussing about these search types, lets see what all options do we have to represent a graph.

------------------------------------------------------------------------------------------------------------------------
Representation of Graph
------------------------------------------------------------------------------------------------------------------------
 * There are two popular representation of graph.

 * Consider a graph:

 *         A------B
 *         |      |
 *         |      |
 *         D------C

 * To store this structure we can use a) Adjacency Matrix, b) Adjacency List

 * Adjacency Matrix representation
 * As the name suggests, using matrix we will represent it. But how will we represent the connection between nodes in
 * the graph? Possibility is each node can be connected to each other node(including itself).
 * What is we keep the number of rows and columns same as the number of nodes in the graph?
 * In this way we would be able to represent connection between nodes, lets say we represent the connection by an int
 * 1 and the non-connections can be represented by 0.
 * The above graph in adjacency matrix representation would look like:

 *     | A | B | C | D |
 *   A | 0 | 1 | 0 | 1 |
 *   B | 1 | 0 | 1 | 0 |
 *   C | 0 | 1 | 0 | 1 |
 *   D | 1 | 0 | 1 | 0 |

 * 1 in the matrix denotes that node represented by row and the node represented by column are connected by an edge in
 * the graph.
 * In case we have a weighted graph, then instead of writing 1 or 0, we will write the edge weight in the matrix.

 * Adjacency List representation:
 * For each node in the graph we can have a list. The list length will be equal to the number of nodes which are
 * connected to it.
 * Eg: representation of above graph would look like:

     | A |  |----->|B| |----->|D| |
     | B |  |----->|A| |----->|C| |
     | C |  |----->|B| |----->|D| |
     | D |  |----->|A| |----->|C| |

 * NOTE: in a directed graph the number of elements present in adjacency list representation will be E.
 * In un-directed graph the number of elements present in adjacency list represetation will be 2*E, reason is every edge
 * will contribute to two connections (a->b and b->a).

 * Which representation to use?
 * It depends on the type of graph. We would want to use minimum space for the representation in our memory.
 * If the graph is dense(i.e. number of edges are very high) then it is better to use matrix.

 * Dense Graph:
     * E = O(V^2), edge is of order vertical(node)*vertical
     * This says graph is almost complete, so matrix will be suited here, coz most of the cells in matrix will be used.
     * Space required: O(n^2), n is number of nodes in graph
 * Sparse Graph:
     * E = O(V), edge is of order of verticals(nodes), there would be as many nodes as the edges in the graph
     * Using Adjacency list would make more sense.
     * Space required: O(N+2*E) or O(N+E), N: number of nodes and E is number of edge

 * Depending on which representation we are going to use will have impact on time complexity of the algorithm on the graph.

------------------------------------------------------------------------------------------------------------------------
Searching methods over graph
------------------------------------------------------------------------------------------------------------------------

 * 1) Breadth First Search
 * 2) Depth First Search

 * In the process of search(visiting nodes) we need to change the state of the nodes(verticals) in the graph as and when
 * they are visited.
 * possible states:
     * visited  : we have visited that node
     * un-visited
     * explored  : we have visited this node as well as we have visited all nodes adjacent to it

 * The idea behind search would be, we start off with DFS or BFS with every node which is unvisited or unexplored and
 * by the time we end the algorithm we are suppose to visit and explore all the nodes.

 * There will be 3 stages involved for every node in the graph:
 * 1) node is not visited and node is not explored
 * 2) node is visited but node is not explored
 * 3) node is visited and node is explored

How to know if node is visited or not?
 * In order to know if a node in graph is visited or not, in our algorithm we can maintain an map, each key in the map
 * will represent a node in our graph. Value of that key can be true or false, depending on if node is visited or not.
 * Let's name this map as visited_map

How to know if a node is explored or not?
 * We can either use a queue or a stack, we will add item to them and then pop items from them, as soon as the queue or
 * the stack gets empty, we would be sure that we have explored all the nodes in the graph.

 * If we use queue in our algorithm, then we would know that items coming out of queue will come in the same order as
 * they were inserted.
 * Now think on which nodes will be inserted first into the queue?
 * The first inserted nodes will definitely be the adjacent nodes of the node from where we have started exploring.
 * So, while using queue we would perform BFS.

 * If we use stack to keep track of unexplored nodes, we would end up performing DFS, as the last inserted node
 * will get popped first and exploration will happen that way.

 * Note that while doing level order traversal, we never maintained an auxiliary map, like we will be using here the
 * visited_map, the reason is in tree we were sure that once a node is visited, it will never get visited again. But
 * with graph there can be several in-degree edges to same node.

 * Below is code for BFS, time complexity:
 - According to the algorithm, each node in connected segment of graph will be added to queue and we are iterating
   over the list which is present in our adjacency map where key is element popped out of queue, so we are executing
   the same time as the number of items are present in our adjacency list representation.
   Thus, TC: O(V+E), V: number of nodes, E: is number of edges, if graph is undirected the time complexity would be
   TC: O(V+2E), which is same as O(V+E)

 * If we use matrix for representation of graph, then the time complexity would become O(N^2), N is number of nodes, coz
 * I will need to iterate over each item in the matrix.

NOTE END */

// adjacency list representation of graph
type Graph struct {
	nodes map[int][]int
}

func (g *Graph) AddEdge(key, destVal int) {
	// check if key is present or not
	if val, ok := g.nodes[key]; ok {
		val = append(val, destVal)
		g.nodes[key] = val
	} else {
		g.nodes[key] = []int{destVal}
	}
}

func (g *Graph) BFS(startKey int) []int {

	res := make([]int, 0)  // here I will store result of search

	// to keep track of nodes for visited state
	visitedMap := make(map[int]bool)

	// to keep track of nodes for exploration state
	queue := make([]int, 0)
	queue = append(queue, startKey)

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]  // popping item from queue, coz we have already printed the value
		res = append(res, item)
		visitedMap[item] = true
		for _, val := range g.nodes[item] {
			if v, ok := visitedMap[val]; !ok || !v {
				// add to queue
				queue = append(queue, val)
				visitedMap[val] = true
			}
		}
	}
	return res
}

// this method will mutate the map passed in the argument
func (g *Graph) bfs(visitedMap map[int]bool, key int) []int {
	res := make([]int, 0)

	queue := make([]int, 0)
	queue = append(queue, key)
	visitedMap[key] = true
	for len(queue) > 0 {
		topElement := queue[0]
		queue = queue[1:]
		res = append(res, topElement)
		for _, val := range g.nodes[topElement] {
			if v, ok := visitedMap[val]; !ok || !v {
				queue = append(queue, val)
				visitedMap[val] = true
			}
		}
	}
	return res
}

// time complexity: O(V+E)
func (g *Graph) BFT(startKey int) []int {
	res := make([]int, 0)
	visitedMap := make(map[int]bool)
	// initializing the visitedMap for all the nodes present in graph
	for k, _ := range g.nodes {
		visitedMap[k] = false
	}
	res = append(res, g.bfs(visitedMap, startKey)...)

	for k, _ := range visitedMap {
		if !visitedMap[k] {
			res = append(res, g.bfs(visitedMap, k)...)
		}
	}
	return res
}

func (g *Graph) DFS(startKey int) []int {
	res := make([]int, 0)
	visitedMap := make(map[int]bool)
	g.dfs(startKey, visitedMap, &res)
	return res
}

func (g *Graph) DFT(startKey int) []int {
	visitedMap := make(map[int]bool)
	res := make([]int, 0)
	// initialize visitedMap with all the nodes available in graph
	for key, _ := range g.nodes {
		visitedMap[key] = false
	}
	g.dfs(startKey, visitedMap, &res)
	for k, _ := range visitedMap {
		if !visitedMap[k] {
			g.dfs(k, visitedMap, &res)
		}
	}
	return res
}

func (g *Graph) dfs(key int, visitedMap map[int]bool, res *[]int) {
	*res = append(*res, key)
	visitedMap[key] = true

	for _, nodeVal := range g.nodes[key] {
		if val, ok := visitedMap[nodeVal]; !ok || !val {
			g.dfs(nodeVal, visitedMap, res)
		}
	}
}

func Graph_basics() string {

	graph1 := GetDummyGraph1()
	bfsRes := graph1.BFS(2)
	result := "BFS result: "
	result += strings.Join(intSliceToStrSlice(bfsRes), ",")
	result += "\nBFT result: "
	bftRes := graph1.BFT(2)
	result += strings.Join(intSliceToStrSlice(bftRes), ",")

	graph2 := GetDummyGraph2()
	dfsRes := graph2.DFS(1)
	result += "\n\nDFS Result:"
	result += strings.Join(intSliceToStrSlice(dfsRes), ",")
	dftRes := graph2.DFT(1)
	result += "\nDFT Result:"
	result += strings.Join(intSliceToStrSlice(dftRes), ",")
	return result
}


/*
    0------->1
   ^|      /
   ||     /                    8----->9
   ||    /                     |     /
   ||   /                      |    /
   ||  /                       |   /
   |v v         |---|          v v
    2---------->3 <-|          11
*/
func GetDummyGraph1() *Graph {
	graph := &Graph{nodes: map[int][]int{}}
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 0)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 3)

	graph.AddEdge(8, 9)
	graph.AddEdge(8, 11)
	graph.AddEdge(9, 11)
	return graph
}


/*
             1
           /  \
          2    3
        /  \  /  \
       4   5 6   7      20----22
        \  | |  /
         \ | | /
            8
 */
func GetDummyGraph2() *Graph {
	graph := &Graph{nodes: map[int][]int{}}
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(3, 6)
	graph.AddEdge(3, 7)
	graph.AddEdge(4, 8)
	graph.AddEdge(5, 8)
	graph.AddEdge(6, 8)
	graph.AddEdge(7, 8)

	graph.AddEdge(20, 22)
	return graph
}

func intSliceToStrSlice(arr []int) []string {
	res := make([]string, 0)
	for _, v := range arr {
		res = append(res, strconv.Itoa(v))
	}
	return res
}

/*
    0------->1
   ^|      /
   ||     /                    8----->9
   ||    /                     |     /
   ||   /                      |    /
   ||  /                       |   /
   |v v         |---|          v v
    2---------->3 <-|          11

     0  1  2  3  4  5  6
  0  0  1  1  0  0  0  0
  1  0  0  1  0  0  0  0
  2  1  0  0  1  0  0  0
  3  0  0  0  1  0  0  0
  4  0  0  0  0  0  1  1
  5  0  0  0  0  0  0  1
  6  0  0  0  0  0  0  0
*/
func GetDummyGraphMatrix1() ([][]int, map[int]int) {
	nodeIdToValueMap := make(map[int]int)
	nodeIdToValueMap[0] = 0
	nodeIdToValueMap[1] = 1
	nodeIdToValueMap[2] = 2
	nodeIdToValueMap[3] = 3
	nodeIdToValueMap[4] = 8
	nodeIdToValueMap[5] = 9
	nodeIdToValueMap[6] = 11

	graph := make([][]int, 7)
	graph[0] = []int{0,  1,  1,  0,  0,  0,  0}
	graph[1] = []int{0,  0,  1,  0,  0,  0,  0}
	graph[2] = []int{1,  0,  0,  1,  0,  0,  0}
	graph[3] = []int{0,  0,  0,  1,  0,  0,  0}
	graph[4] = []int{0,  0,  0,  0,  0,  1,  1}
	graph[5] = []int{0,  0,  0,  0,  0,  0,  1}
	graph[6] = []int{0,  0,  0,  0,  0,  0,  0}
	return graph, nodeIdToValueMap
}
