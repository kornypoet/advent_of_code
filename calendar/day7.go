package calendar

import (
	"io"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/kornypoet/advent_of_code/util"
)

type Node struct {
	location string
	parent   *Node
	children []*Node
	files    []*File
}

type File struct {
	name string
	size int
}

type FileTree struct {
	nodes   []*Node
	current *Node
}

func (ft *FileTree) display() {
	var root *Node
	for _, node := range ft.nodes {
		if node.location == "/" {
			root = node
		}
	}
	root.display(0, "")
}

func (n *Node) display(indent int, prefix string) {
	log.Printf("%s%s%s %d", strings.Repeat(" ", indent), prefix, n.location, n.size())
	for _, ch := range n.children {
		ch.display(indent+2, prefix+n.location+"/")
	}
	for _, f := range n.files {
		f.display(indent + 2)
	}
}

func (f *File) display(indent int) {
	log.Printf("%s%s %d", strings.Repeat(" ", indent), f.name, f.size)
}

func (ft *FileTree) add(node *Node) {
	log.Printf("current node %#v", ft.current)
	if ft.current == nil {
		ft.nodes = append(ft.nodes, node)
		ft.current = node
	} else {
		// var exists bool
		// for _, n := range cur.children {
		//	if n == node {
		//		exists = true
		//		break
		//	}
		// }
		// if exists {
		//	// node already exists in tree
		//	ft.current = node
		// } else {
		ft.current.add(node)
		node.parent = ft.current
		ft.nodes = append(ft.nodes, node)
		ft.current = node
		// }
	}
}

func (ft *FileTree) backup() {
	cur := ft.current
	ft.current = cur.parent
}

func (n *Node) add(node *Node) {
	n.children = append(n.children, node)
}

func (n *Node) addFile(f *File) {
	n.files = append(n.files, f)
}

func (n *Node) size() (total int) {
	for _, f := range n.files {
		total += f.size
	}
	for _, n := range n.children {
		total += n.size()
	}
	return
}

func Day7(input io.Reader, part int) int {
	var total int
	var tree = &FileTree{}
	util.ProcessByLine(input, func(line string, num int) {
		cmdRxp := regexp.MustCompile(`\$ cd ([\/\.\w]+)`)
		command := cmdRxp.FindStringSubmatch(line)
		lsRxp := regexp.MustCompile(`\$ ls`)
		listItems := lsRxp.FindStringSubmatch(line)
		if len(command) != 0 {
			log.Print(command[0])
			log.Printf("Changing directory to %s", command[1])
			switch command[1] {
			case "..":
				tree.backup()
			default:
				node := &Node{location: command[1]}
				tree.add(node)
			}
		} else if len(listItems) != 0 {
			log.Print("listing items")
		} else {
			contentRxp := regexp.MustCompile(`(dir|[\d]+) ([\.\w]+)`)
			contents := contentRxp.FindStringSubmatch(line)
			if len(contents) == 0 {
				log.Fatalf("line didn't parse %s", line)
			} else {
				switch contents[1] {
				case "dir":
					node := &Node{location: contents[2]}
					tree.current.add(node)
				default:
					filesize, _ := strconv.Atoi(contents[1])
					filename := contents[2]
					tree.current.addFile(&File{size: filesize, name: filename})
				}
			}
		}
	})

	if part == 1 {
		for _, n := range tree.nodes {
			log.Printf("%s %d", n.location, n.size())
			if n.size() <= 100000 {
				total += n.size()
			}
		}
	} else {
		sizes := []int{}
		for _, n := range tree.nodes {
			log.Printf("%s %d", n.location, n.size())
			if n.size() >= 528671 {
				sizes = append(sizes, n.size())
			}
		}
		sort.Ints(sizes)
		log.Print(sizes[0])
		// rootFs := sizes[len(sizes) -1]
		// totalFs := 70000000
		// unusedFs := totalFs - rootFs
		// spaceNeeded := 30000000 - unusedFs
		// log.Print(spaceNeeded)

	}
	tree.display()
	log.Printf("Total is %d", total)
	return total
}
