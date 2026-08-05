package tree

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mysticCoder100/algorithm/graph"
)

/**
 * A simplpe traversal of a directory
 */
func PrintFileName(path string) {
	myQueue := graph.MyQueue[string]{}
	myQueue.Enqueue(path)

	for !myQueue.IsEmpty() {
		dir, _ := myQueue.Dequeue()
		files, err := os.ReadDir(dir)

		if err != nil {
			fmt.Println("Unable to read directory:", dir)
			continue // Use continue so one bad read doesn't kill the whole program
		}

		for _, file := range files {
			fullPath := filepath.Join(dir, file.Name())

			if file.IsDir() {
				fmt.Println(file.Name(), "[Subdirectory -> Enqueued]")
				myQueue.Enqueue(fullPath)
				continue
			}
			fmt.Println("File:", fullPath)
		}
	}
}

/**
 * A depth first version
 */
func DepthFirstSearch(path string) {
	entries, err := os.ReadDir(path)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range entries {
		fullPath := filepath.Join(path, file.Name())

		if file.IsDir() {
			DepthFirstSearch(fullPath)
			continue
		}

		fmt.Println(fullPath)
	}
}
