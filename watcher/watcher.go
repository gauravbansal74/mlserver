package watcher

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/gauravbansal74/mlserver/pkg/logger"
	"github.com/gauravbansal74/mlserver/pkg/rabbitmq"
	"github.com/radovskyb/watcher"
)

const (
	createFile = "CREATE"
)

// FileMsg for rabbitmq
type FileMsg struct {
	Name string `json:"name"`
}

// Init watcher init function
func Init(folderName string) {
	w := watcher.New()
	w.IgnoreHiddenFiles(true)
	// Watch this folder for changes.
	if err := w.Add(folderName); err != nil {
		log.Println(err)
	}
	go func() {
		for {
			select {
			case event := <-w.Event:
				if !event.IsDir() {
					fileName := event.Name()
					if event.Op.String() == createFile && strings.HasSuffix(fileName, ".csv") {

						fileMsg, err := json.Marshal(FileMsg{
							Name: fileName,
						})
						if err != nil {
							logger.Error(err, "error while creating msg for rabbitmq")
						}
						err = rabbitmq.Publish(fileMsg)
						if err != nil {
							logger.Error(err, "error while pushing message into rabbitMQ")
						}
					}
				}
			case err := <-w.Error:
				log.Println(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// Print a list of all of the files and folders currently
	// being watched and their paths.
	for _, f := range w.WatchedFiles() {
		if !f.IsDir() {
			logger.Info("watched file: " + f.Name())
		}
	}

	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Println(err)
	}
}
