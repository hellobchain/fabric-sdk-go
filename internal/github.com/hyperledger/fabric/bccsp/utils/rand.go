package utils

import (
	"github.com/wsw365904/wswlog/wlogging"
	"sync"

	"io"
)

var randlogger = wlogging.MustGetLoggerWithoutName()
var randReader *pipelineRandReader

var randRDPool *sync.Pool
var randReaderOnce sync.Once
var rand_cache_baselen int

const (
	rand_once_baselen = 16
	defaut_cache_size = 12800
)

func NewRandReader(randCacheSize int, randBytesLen int, reader io.Reader) io.Reader {
	randReaderOnce.Do(func() {
		if randBytesLen <= 0 {
			rand_cache_baselen = rand_once_baselen
		} else {
			rand_cache_baselen = randBytesLen
		}

		if randCacheSize <= 0 {
			randCacheSize = defaut_cache_size
		}

		randRDPool = &sync.Pool{New: func() interface{} {
			return make([]byte, rand_cache_baselen)
		}}

		randReader = &pipelineRandReader{
			randcache: make(chan []byte, randCacheSize),
			randbytesPool: &sync.Pool{New: func() interface{} {
				return make([]byte, rand_cache_baselen)
			}},
			rd: reader,
		}
		randReader.start()
	})
	return randReader
}

type pipelineRandReader struct {
	randbytesPool *sync.Pool
	randcache     chan []byte
	rd            io.Reader
}

func (reader *pipelineRandReader) Read(b []byte) (n int, err error) {

	cache := <-reader.randcache
	rdlen := copy(b, cache)
	reader.randbytesPool.Put(cache)
	return rdlen, nil
}
func (reader *pipelineRandReader) start() {
	go func() {
		for {
			rdbuf := reader.randbytesPool.Get().([]byte)
			reader.rd.Read(rdbuf)
			reader.randcache <- rdbuf
		}
	}()

}
