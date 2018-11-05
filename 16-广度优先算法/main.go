
package main 

import(
	"fmt"
	"os"
)

// 迷宫二维数组
func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	// 二维数组/slice
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze

}

// 坐标 结构体
type point struct {
	i, j int
}

// 路线 四个方向
var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

// 算下一步的坐标
func (p point) add(r point) point{
	// 返回一个新的point，不使用*point改变原point
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	// 横向越界
	if p.i < 0 || p.i >= len(grid) {
		fmt.Println(p.i,"yue")
		return 0, false
	}
	// // 纵向越界
	if p.j < 0 || p.j >=  len(grid[p.i]) {
		fmt.Println(p.i,"gou")
		return 0, false
	}

	return grid[p.i][p.j], true
}
// 走迷宫
func walk(maze [][]int, start, end point) [][]int{
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	// 坐标队列
	que := []point{start}

	for len(que) > 0{
		cur := que[0]
		que = que[1:]

		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)
			// 探索下一个节点（能走）
			// 下一步是0 next is 0
			// 下一步不是起点 next != start

			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}

			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1

			que = append(que, next)
		}
	}

	return steps
}
func main() {
	maze := readMaze("./maze.in")

	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%d ",val)
		}
		fmt.Println()
	}
	fmt.Println()
	steps := walk(maze, point{0, 0}, 
		point{len(maze) - 1, len(maze[0]) - 1})

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%d ",val)
		}
		fmt.Println()
	}
}

















