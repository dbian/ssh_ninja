package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"

	"gopkg.in/yaml.v3"
)

type Cfg struct {
	Hosts []string `yaml:"hosts"`
	Cmd   string   `yaml:"cmd"`
}

func main() {
	var cmd *exec.Cmd
	cfg := Cfg{}

	dat, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(dat, &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("cfg:\n%v\n\n", cfg)
	wg := sync.WaitGroup{}
	for _, v := range cfg.Hosts {
		wg.Add(1)
		go func(host string) {
			cmd = exec.Command("ssh", host, cfg.Cmd)
			if whoami, err := cmd.Output(); err != nil {
				fmt.Println("hosts execute err:", err)
			} else {
				fmt.Println("hosts execute done", string(whoami))
			}
			wg.Done()
		}(v)
	}
	wg.Wait()

	fmt.Println("bye.")

}
