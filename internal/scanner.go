package internal

import (
	"math"
	"os"
	"runtime"
	"sync"
	"time"

	// log "github.com/sirupsen/logrus"

	"github.com/daan-hoogland/walk"
)

//File struct used to store file information.
type File struct {
	info os.FileInfo
	path string
	err  error
}

type resultList struct {
	fileMatches []File
	sync.RWMutex
}

var (
	wg sync.WaitGroup
	//Results is a thread safe list containing results of the scanner action.
	Results resultList

	exp *Expected

	fileQueue = make(chan File, 300)
)

//StartJobs starts the consumers and producer.
func StartJobs(expected *Expected, maxProcs int, rootDir string) {
	exp = expected
	procs := calcProcs(maxProcs)
	consumers := procs[1]
	if !(consumers-1 > 0) {
		consumers = 1
	}

	for i := 0; i < consumers; i++ {
		// log.WithField("component", "producer").Traceln("starting producer " + strconv.Itoa(i))
		wg.Add(1)
		go consume(i)
	}

	runtime.GOMAXPROCS(maxProcs)
	wg.Add(1)
	go scan(rootDir, procs[0])

	wg.Wait()
}

func consume(id int) {
	defer wg.Done()
	for file := range fileQueue {
		// log.WithField("component", "new job").Debugln("consumer " + strconv.Itoa(id))
		res := MatchFile(file.info, file.path, exp)
		if res.ValidateResult(exp) {
			addResult(file)
		}
	}
}

func addResult(file File) {
	Results.Lock()
	defer Results.Unlock()
	Results.fileMatches = append(Results.fileMatches, file)
}

func addFileToQueue(file File) {
	select {
	case fileQueue <- file:
		// log.Debugln("adding file")
	default:
		// log.Debugln("queue full, waiting")
		// log.Debugln(len(fileQueue))
		time.Sleep(50 * time.Millisecond)
		// log.Debugln(len(fileQueue))
		addFileToQueue(file)
	}
}

//Scan starts the filesystem scanning
func scan(directory string, maxProcs int) {
	defer wg.Done()
	err := walk.CustomWalk(directory, visit, int(math.Ceil(0.2*float64(maxProcs))))
	close(fileQueue)
	if err != nil {
		//throw big error
	}
}

func visit(path string, info os.FileInfo, err error) error {
	if !info.IsDir() {
		addFileToQueue(File{
			info: info,
			path: path,
			err:  err,
		})
	}
	return nil
}

func calcProcs(maxProcs int) []int {
	walkProcs := int(math.Ceil(0.2 * float64(maxProcs)))
	analyseProcs := maxProcs - walkProcs
	return []int{walkProcs, analyseProcs}
}
