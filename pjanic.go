package miralempjanic

import (
	"encoding/json"
	"fmt"
	"github.com/go-errors/errors"
	"io/ioutil"
	"os"
)

type Miralem struct {
	Path         string `json:"-"`
	CurrentJobID string `json:"current_job"`
	Jobs         []Job  `json:"jobs"`
}

type Job struct {
	ID            string `json:"current_job"`
	LastProcessID string `json:"last_process_id"`
	Done          bool   `json:"done"`
}

func New(path string) *Miralem {
	return &Miralem{
		Path: path,
	}
}

func FromFile(path string) (mir Miralem) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(file, &mir)
	mir.Path = path

	return
}

func (self *Miralem) RegisterJob(id string) {
	self.Jobs = append(self.Jobs, Job{ID: id})
}

func (self *Miralem) CloseJob(id string) {
	for _, job := range self.Jobs {
		if job.ID == id {
			job.Done = true
		}
	}
}

func (self *Miralem) serialize() {
	serialization, _ := json.Marshal(self)
	ioutil.WriteFile(self.Path, serialization, 0644)
}

func (self *Miralem) Pjanic(processID string, err error) {
	fmt.Println("Panic on the streets of London")
	fmt.Println("Panic on the streets of Birmingham")
	fmt.Println("I wonder to myself")
	fmt.Println("Could life ever be sane again?")
	err = errors.New(errors.Errorf(err.Error()))
	fmt.Println(err.(*errors.Error).ErrorStack())
	self.serialize()
	os.Exit(1)
}
