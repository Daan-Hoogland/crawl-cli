package internal

import (
	"math"
	"os"
	"runtime"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/daan-hoogland/walk"
)

//File struct used to store file information.
type File struct {
	info os.FileInfo
	path string
	err  error
}

type resultList struct {
	FileMatches []File
	sync.RWMutex
}

var (
	wg sync.WaitGroup
	//Results is a thread safe list containing results of the scanner action.
	Results resultList

	fileQueue = make(chan File, 300)
)

//StartJobs starts the consumers and producer.
func StartJobs() {
	// log.WithField("component", "producer").Traceln("entering start jobs")
	consumers := MaxProcs - int(math.Ceil(0.2*float64(MaxProcs)))
	if !(consumers-1 > 0) {
		consumers = 1
	}

	for i := 0; i < consumers; i++ {
		// log.WithField("component", "producer").Traceln("starting producer " + strconv.Itoa(i))
		wg.Add(1)
		go consume(i)
	}

	runtime.GOMAXPROCS(MaxProcs)
	wg.Add(1)
	go scan(Directory)

	wg.Wait()
}

func consume(id int) {
	defer wg.Done()
	for file := range fileQueue {
		// log.WithField("component", "new job").Debugln("consumer " + strconv.Itoa(id))
		res := MatchFile(file.info, file.path)
		if ValidateResult(res) {
			addResult(file)
		}
	}
}

func addResult(file File) {
	Results.Lock()
	defer Results.Unlock()
	Results.FileMatches = append(Results.FileMatches, file)
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
func scan(directory string) {
	defer wg.Done()
	err := walk.CustomWalk(directory, visit, int(math.Ceil(0.2*float64(MaxProcs))))
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
