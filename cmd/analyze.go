package cmd

import (
	"lid/models"
	"lid/services/disk_inventory"
	"lid/services/file_ops"
	"lid/services/logger"
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/cobra"
)

var save bool
var refStorageDeviceUUID string

func analyzeNode(log *logger.LoggerInstance, wg *sync.WaitGroup, in <-chan string, out chan<- models.File, errored chan<- struct{}) {
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
		out <- models.File{Name: filepath.Base(filename), Size: fileStat.Size(), MD5: md5Hash}
		f.Close()
	}
}

// analyzeCmd represents the analyze command
var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze a set of directories and/or files",
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.CreateLogger("analyze")
		var fl models.FileList
		var refStorageDevice models.Device

		log.Trace("analyze file called")

		if save {
			var err error
			refStorageDevice, err = drs.FindByUUID(refStorageDeviceUUID)
			if err != nil {
				log.Fatal("device `"+refStorageDeviceUUID+"` couldn't be loaded from the db", err)
				return
			}
		}

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
		supplier, receiver, errored, done := make(chan string), make(chan models.File), make(chan struct{}), make(chan struct{})
		received := 0

		for i := 0; i < numRoutines; i++ {
			wg.Add(1)
			go analyzeNode(&log, &wg, supplier, receiver, errored)
		}

		go func() {
		OUT:
			for {
				select {
				case f := <-receiver:
					fl.Append(f)
					if save {
						frs.Create(f.Name, f.MD5, f.Size, refStorageDevice)
					}
					received++
				case <-errored:
					received++
				}

				if received == len(files) {
					done <- struct{}{}
					break OUT
				}
			}
		}()

		for _, filename := range files {
			supplier <- filename
		}

		<-done
		close(supplier)
		close(receiver)
		close(errored)
		close(done)
		wg.Wait()

		if tabulate {
			fl.Pretty()
		}
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
	analyzeCmd.PersistentFlags().BoolP("tabulate", "t", false, "Tabulate analysis")
	analyzeCmd.Flags().IntP("routines", "r", 1, "number of go routines to use while hashing")
	analyzeCmd.Flags().BoolVar(&save, "save", false, "save analyzed files to db")
	analyzeCmd.Flags().StringVar(&refStorageDeviceUUID, "device-uuid", "", "associate analyzed files to this storage device by uuid")
	analyzeCmd.MarkFlagsRequiredTogether("save", "device-uuid")
}
