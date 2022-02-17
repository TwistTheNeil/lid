package cmd

import (
	"lid/models"
	"lid/services/disk_inventory"
	"lid/services/file_ops"
	"lid/services/logger"
	"os"
	"sync"

	"github.com/spf13/cobra"
)

func analyzeNode(log *logger.LoggerInstance, wg *sync.WaitGroup, in <-chan string, out chan<- models.Node, errored chan<- struct{}) {
	defer wg.Done()

	for filename := range in {
		f, err := os.Open(filename)
		if err != nil {
			log.Error("something went wrong opening file", err, "file", filename)
			errored <- struct{}{}
			continue
		}

		fileStat, err := f.Stat()
		if err != nil {
			log.Error("Can't read file", err, "file", filename)
			errored <- struct{}{}
			continue
		}

		md5Hash := file_ops.MD5Hash(f)
		out <- models.Node{Name: filename, Size: fileStat.Size(), MD5: md5Hash}
		f.Close()
	}
}

// analyzeCmd represents the analyze command
var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze a set of directories and/or files",
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.CreateLogger("analyze")
		var nl models.NodeList

		log.Trace("analyze file called")

		tabulate, err := cmd.Flags().GetBool("tabulate")
		if err != nil {
			log.Error("Can't read --tabulate value", err)
		}

		numRoutines, err := cmd.Flags().GetInt("routines")
		if err != nil {
			log.Error("Can't get numRoutines. Defaulting to 1", err)
			numRoutines = 1
		}

		files := disk_inventory.BuildFileList(args...)
		var wg sync.WaitGroup
		supplier, receiver, errored, done := make(chan string), make(chan models.Node), make(chan struct{}), make(chan struct{})
		received := 0

		for i := 0; i < numRoutines; i++ {
			wg.Add(1)
			go analyzeNode(&log, &wg, supplier, receiver, errored)
		}

		go func() {
			for {
				select {
				case n := <-receiver:
					nl.Append(n)
					received++
				case <-errored:
					received++
				}

				if received == len(files) {
					done <- struct{}{}
					break
				}
			}
		}()

		for _, filename := range files {
			supplier <- filename
		}

		<-done
		close(supplier)
		close(receiver)
		wg.Wait()

		if tabulate {
			nl.Pretty()
		}
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
	analyzeCmd.PersistentFlags().BoolP("tabulate", "t", false, "Tabulate analysis")
	analyzeCmd.Flags().IntP("routines", "r", 1, "number of go routines to use while hashing")
}
