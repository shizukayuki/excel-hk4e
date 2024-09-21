package excel

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var resources = map[string]any{}

type LoaderConfig struct {
	Root     string
	ReadFile func(root, name string) ([]byte, error)
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
	for name, v := range resources {
		if err := cfg.readFile(name, v); err != nil {
			return fmt.Errorf("failed to load %s: %w", name, err)
		}
	}
	return nil
}
