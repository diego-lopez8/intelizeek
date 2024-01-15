package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path"
	"time"
)

var syslogWriter *syslog.Writer

func main() {
	// initialize logger
	var err error
	syslogWriter, err = syslog.New(syslog.LOG_INFO, "intellizeek")
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(syslogWriter)
	defer syslogWriter.Close()

	// flags
	trainMode := flag.Bool("train", false, "Enable training mode")
	inferenceMode := flag.Bool("inference", false, "Enable inference mode")
	logDir := flag.String("logdir", "/opt/zeek/logs", "Directory for log files")
	flag.Parse()

	if !*trainMode && !*inferenceMode {
		fmt.Println("Please specify a mode: -train or -inference")
		flag.Usage()
		os.Exit(1)
	}

	if *trainMode {
		log.Println("Starting in Training mode")
		train(*logDir)
		return
	}
	if *inferenceMode {
		log.Println("Starting in Inference mode")
		inference(*logDir)
		return
	}
}

func train(logDir string) {
	// TODO Write training code
	return
}

func inference(logDir string) {
	log.Printf("Inference on logs in %s\n", logDir)
	logFilePath := path.Join(logDir, "current/conn.log")

	// handle file I/O
	file, err := os.Open(logFilePath)
	if err != nil {
		log.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	// Loop for file read
	// TODO write inference code on line
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			// sleep for 1s
			time.Sleep(1 * time.Second)
			continue
		}
		fmt.Print(line)
	}
}
