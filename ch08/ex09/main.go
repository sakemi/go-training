package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

type rootInfo struct {
	sizes  chan int64
	nbytes int64
	nfiles int64
	end    bool
	n      *sync.WaitGroup
}

func main() {
	flag.Parse()

	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := map[string]*rootInfo{}
	for _, root := range roots {
		root := root
		fileSizes[root] = &rootInfo{make(chan int64), 0, 0, false, &sync.WaitGroup{}}
		fileSizes[root].n.Add(1)
		go func() {
			fileSizes[root].n.Wait()
			close(fileSizes[root].sizes)
		}()
		go walkDir(root, fileSizes[root].n, fileSizes[root].sizes)
	}

	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(50 * time.Millisecond)
	}
	var nfiles, nbytes int64
	endCnt := 0

	for {
		if endCnt == len(fileSizes) {
			break
		}
		for k, v := range fileSizes {
			if v.end {
				continue
			}
			select {
			case size, ok := <-v.sizes:
				if !ok {
					endCnt++
					v.end = true
				}
				nfiles++
				nbytes += size
				v.nfiles++
				v.nbytes += size
			case <-tick:
				printDirDiskUsage(k, v.nfiles, v.nbytes)
			}
		}
	}

	printDiskUsage(nfiles, nbytes)
}

func printDirDiskUsage(dir string, nfiles, nbytes int64) {
	fmt.Printf("%v: %d files  %.1f GB\n", dir, nfiles, float64(nbytes)/1e9)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
