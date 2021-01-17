package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

// Configs of the watcher
type Configs struct {
	Setup    []string `json:"setup"`
	Build    string   `json:"build"`
	Run      string   `json:"run"`
	Cleanup  []string `json:"cleanup"`
	Excludes []string `json:"excludes"`
}

func load(filename string) (c Configs, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()

	if err = json.NewDecoder(f).Decode(&c); err != nil {
		return
	}
	return
}

// Exclude path ?
func (c Configs) Exclude(path string, info os.FileInfo) bool {
	for _, exclude := range c.Excludes {
		if strings.HasPrefix(path, exclude) {
			return true
		}
	}
	return false
}

func exit(f string, a ...interface{}) {
	if !strings.HasSuffix(f, "\n") {
		f += "\n"
	}
	fmt.Fprintf(os.Stderr, "ERROR: %s", fmt.Sprintf(f, a...))
	os.Exit(1)
}

func info(f string, a ...interface{}) {
	if !strings.HasSuffix(f, "\n") {
		f += "\n"
	}
	fmt.Fprintf(os.Stderr, "INFO : %s", fmt.Sprintf(f, a...))
}

// Watcher ...
type Watcher struct {
	watcher *fsnotify.Watcher
	conf    Configs
}

// NewWatcher creates a new Watcher instance
func NewWatcher(conf Configs) (*Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	return &Watcher{
		watcher: watcher,
		conf:    conf,
	}, nil
}

// Add directory to watching list
func (w *Watcher) Add(dir string) error {
	return filepath.Walk(dir, func(path string, finfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !finfo.IsDir() || w.conf.Exclude(path, finfo) {
			return nil
		}
		info("- watch %s", path)
		return w.watcher.Add(path)
	})
}

// Remove ...
func (w *Watcher) Remove(dir string) error {
	return filepath.Walk(dir, func(path string, finfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !finfo.IsDir() || w.conf.Exclude(path, finfo) {
			return nil
		}
		info("- unwatch %s", path)
		return w.watcher.Remove(path)
	})
}

// Close ...
func (w *Watcher) Close() error { return w.watcher.Close() }

// Events ...
func (w *Watcher) Events() <-chan fsnotify.Event { return w.watcher.Events }

// Errors ...
func (w *Watcher) Errors() <-chan error { return w.watcher.Errors }

// Runner ...
type Runner struct {
	conf Configs
	prev *exec.Cmd
}

// NewRunner creates a new Runner instance
func NewRunner(conf Configs) (*Runner, error) {
	return &Runner{
		conf: conf,
		prev: nil,
	}, nil
}

// Run ...
func (r *Runner) Run() error {
	if r.prev != nil {
		info("STEP pre-build: start")
		info("STEP pre-build: kill pid=%d", r.prev.Process.Pid)
		if err := r.prev.Process.Kill(); err != nil {
			info("ERROR: STEP pre-build: kill prev run (pid=%d): %v", r.prev.Process.Pid, err)
		} else if _, err := r.prev.Process.Wait(); err != nil {
			info("ERROR: STEP pre-build: wait for prev run command exit: %v", err)
		}
	}
	info("STEP build: start")
	build := exec.Command("/bin/sh", "-c", r.conf.Build)
	build.Stdout = os.Stdout
	build.Stderr = os.Stderr
	if err := build.Run(); err != nil {
		return fmt.Errorf("STEP build: %v", err)
	}
	info("STEP build: finished")
	info("STEP run: start")
	run := exec.Command("/bin/sh", "-c", r.conf.Run)
	run.Stdout = os.Stdout
	run.Stderr = os.Stderr
	if err := run.Start(); err != nil {
		return fmt.Errorf("STEP run: %v", err)
	}
	info("STEP run: pid=%d", run.Process.Pid)
	r.prev = run
	return nil
}

// Close ...
func (r *Runner) Close() error {
	if r.prev != nil {
		if err := r.prev.Process.Kill(); err != nil {
			return err
		}
		if _, err := r.prev.Process.Wait(); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	var (
		conffile string
	)
	flag.StringVar(&conffile, "config", "watcher.config.json", "relative path to config file")
	flag.Parse()

	conf, err := load(conffile)
	if err != nil {
		exit("load config: %v", err)
	}
	info("conf: %#v", conf)

	watcher, err := NewWatcher(conf)
	if err != nil {
		exit("NewWatcher: error: %v\n", err)
	}
	defer watcher.Close()

	runner, err := NewRunner(conf)
	if err != nil {
		exit("NewRunner: error: %v\n", err)
	}
	defer runner.Close()

	go func() {
		info("run command")
		if err := runner.Run(); err != nil {
			exit("runner.Run: error: %v\n", err)
		}
	}()

	if err := watcher.Add("."); err != nil {
		exit("failed to watch files: %v", err)
	}
	info("watcher setup")

	done := make(chan bool)
	go func() {
		defer close(done)
		info("watching...")

		var startTime = time.Now()

		for {
			select {
			case event, received := <-watcher.Events():
				info("%s: %s", event.Op, event.Name)
				if !received {
					return
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					if finfo, err := os.Stat(event.Name); err == nil && finfo.IsDir() {
						watcher.Add(event.Name)
						info("watch %s", event.Name)
					}
				} else if event.Op&fsnotify.Remove == fsnotify.Remove {
					watcher.Remove(event.Name)
				}
				if time.Now().Before(startTime) {
					break // break from `select`
				}
				go func() {
					if err := runner.Run(); err != nil {
						info("failed to run: %v", err)
					}
				}()

				startTime = time.Now().Add(1 * time.Second)
			case err, received := <-watcher.Errors():
				if !received {
					return
				}
				exit("watcher.Errors: %v", err)
			}
		}
	}()
	<-done
}
