package terminalW

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/grewwc/go_tools/src/containerW"
	"github.com/grewwc/go_tools/src/utilsW"
)

var Once sync.Once

var DefaultExtensions = containerW.NewSet()

func init() {
	DefaultExtensions.AddAll(".py", ".cpp", ".js", ".txt", ".h", ".c", ".tex", ".html",
		".css", ".java", ".go", ".cc")
}

var Extensions = containerW.NewSet()
var CheckExtension bool
var Exclude bool

var NumPrint int64 = 5

var Count int64

// maximum 5000 threads
var maxThreads = make(chan struct{}, 5000)
var Verbose bool
var CountMu sync.Mutex

// how many levels to search
var MaxLevel int32

// this function is the main part
// acts like a framework
func Find(rootDir string, task func(string), wg *sync.WaitGroup, level int32) {
	defer wg.Done()
	if atomic.LoadInt32(&level) > MaxLevel {
		return
	}
	maxThreads <- struct{}{}
	defer func() { <-maxThreads }()
	CountMu.Lock()
	if Count >= NumPrint {
		CountMu.Unlock()
		Once.Do(func() {
			summaryString := fmt.Sprintf("%d matches found\n", Count)
			fmt.Println(strings.Repeat("-", len(summaryString)))
			matches := int64(math.Min(float64(Count), float64(NumPrint)))
			fmt.Printf("%v matches found\n", matches)
		})
		os.Exit(0)
		return
	}
	CountMu.Unlock()
	subs, err := ioutil.ReadDir(rootDir)
	if err != nil {
		if Verbose {
			fmt.Fprintln(os.Stderr, err)
		}
		return
	}

	for _, sub := range subs {
		subName := path.Join(rootDir, sub.Name())
		extName := path.Ext(subName)
		if sub.IsDir() {
			wg.Add(1)
			go Find(subName, task, wg, atomic.AddInt32(&level, 1))
			atomic.AddInt32(&level, -1)
		} else if !utilsW.IsTextFile(subName) {
			continue
		}
		if !CheckExtension {
			task(subName)
		} else if Extensions.Contains(extName) {
			task(subName)
		}
	}
}
