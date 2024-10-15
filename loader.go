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
	langs := map[string]struct{}{}
	for _, lang := range cfg.Languages {
		name := fmt.Sprintf("TextMap/TextMap%s.json", lang)
		langs[strings.ToLower(name)] = struct{}{}
	}

	errCh := make(chan error, len(resources))
	loadFile := func(wg *sync.WaitGroup, name string, v any) {
		defer wg.Done()

		if s := strings.ToLower(name); strings.HasPrefix(s, "textmap") {
			if _, ok := langs[s]; !ok {
				return
			}
		}

		if err := cfg.readFile(name, v); err != nil {
			errCh <- fmt.Errorf("failed to load %s: %w", name, err)
		}
	}

	var wg sync.WaitGroup
	for name, v := range resources {
		wg.Add(1)
		go loadFile(&wg, name, v)
	}
	wg.Wait()
	close(errCh)

	var err error
	for nerr := range errCh {
		err = errors.Join(err, nerr)
	}
	return err
}
