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
	ch := make(chan (FileSize), len(filePath))
	wg := sync.WaitGroup{}

	for _, filePath := range filePath {
		wg.Add(1)
		go GetPageSize(filePath, &wg, ch)
	}

	wg.Wait()
	close(ch)

	for fileSize := range ch {
		result = append(result, fileSize)
	}

	SortFileSize(&result)

	return result

}

func GetPageSize(url string, wg *sync.WaitGroup, ch chan FileSize) error {
	defer wg.Done()

	res, err := http.Get(url)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	ch <- FileSize{
		Url:  url,
		Size: int64(len(body)),
	}

	return nil
}
