package excel

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var resources = map[string]any{}

type LoaderConfig struct {
	Root      string
	ReadFile  func(root, name string) ([]byte, error)
	Languages []string
}

func (cfg *LoaderConfig) readFile(name string, v any) error {
	var data []byte
	var err error
	if cfg.ReadFile != nil {
		data, err = cfg.ReadFile(cfg.Root, name)
	} else {
		data, err = os.ReadFile(filepath.Join(cfg.Root, name))
	}
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func load(name string, v any) {
	if _, ok := resources[name]; ok {
		panic(fmt.Errorf("%s is already loaded", name))
	}
	resources[name] = v
}

func LoadResources(cfg LoaderConfig) error {
	errCh := make(chan error, len(resources))
	var wg sync.WaitGroup
	wg.Add(len(resources))
	for name, v := range resources {
		var f func(wg *sync.WaitGroup, name string, v any, err chan error)
		switch {
		case strings.HasPrefix(strings.ToLower(name), "textmap"):
			f = cfg.fetchTextMap
		default:
			f = cfg.fetchExcel
		}
		go f(&wg, name, v, errCh)
	}
	wg.Wait()
	close(errCh)

	var err error
	for nerr := range errCh {
		err = errors.Join(err, nerr)
	}
	return err
}

func (cfg *LoaderConfig) fetchExcel(wg *sync.WaitGroup, name string, v any, errCh chan error) {
	defer wg.Done()
	if err := cfg.readFile(name, v); err != nil {
		errCh <- fmt.Errorf("failed to load %s: %w", name, err)
	}
}

func (cfg *LoaderConfig) fetchTextMap(wg *sync.WaitGroup, name string, v any, errCh chan error) {
	defer wg.Done()
	merged := v.(map[TextMapHash]string)
	clear(merged)

	var lang string
	for _, v := range cfg.Languages {
		if strings.EqualFold(fmt.Sprintf("TextMap/TextMap%s.json", v), name) {
			lang = v
			break
		}
	}
	if lang == "" {
		return
	}

	const ChunkMin = 2
	var chunks int
	var err error
	for n := 0; n <= ChunkMin; n++ {
		pname := name
		if n > 0 {
			dot := strings.Index(name, ".")
			pname = fmt.Sprintf("%s_%v.%s", name[:dot], n, name[dot+1:])
		}

		var part map[TextMapHash]string
		if nerr := cfg.readFile(pname, &part); nerr != nil {
			err = fmt.Errorf("failed to load %s: %w", pname, nerr)
			if len(merged) > 0 {
				err = nil
			} else if n == 0 {
				continue
			}
			break
		}

		err = nil
		chunks = n
		for k, v := range part {
			merged[k] = v
		}
		if n == 0 {
			break
		}
	}

	if chunks != 0 && chunks < ChunkMin {
		err = fmt.Errorf("failed to load %s: missing %v chunks out of %v", name, ChunkMin-chunks, ChunkMin)
	}
	errCh <- err
}
