package file

import (
	"io"
	"net/http"
	"sort"
	"sync"
)

type FileSize struct {
	Url  string
	Size int64
}

func SortFileSize(result *[]FileSize) {
	sort.Slice(*result, func(i, j int) bool {
		return (*result)[i].Size > (*result)[j].Size
	})
}

func GetSortedFileSize(filePath []string) []FileSize {
	var result = []FileSize{}
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}

	for _, filePath := range filePath {
		wg.Add(1)
		go GetPageSize(filePath, &result, &mutex, &wg)
	}

	wg.Wait()
	SortFileSize(&result)

	return result

}

func GetPageSize(url string, result *[]FileSize, mutex *sync.Mutex, wg *sync.WaitGroup) error {
	defer wg.Done()
	mutex.Lock()

	defer mutex.Unlock()

	res, err := http.Get(url)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	*result = append(*result, FileSize{
		Url:  url,
		Size: int64(len(body)),
	})

	return nil
}
