package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	FilePath1 := "/data/uploads/2017/07/NDOC_magnetic_tape_library.jpg"
	FilePath2 := "/data/uploads/2017/07/NDOC_magnetic_tape_libraryTemp.jpg"
	FilePath3 := "/data/uploads/2017/07/Tape.jpg"

	for {
		err := os.Rename(FilePath1, FilePath2)
		if err != nil {
			fmt.Printf(err.Error())
		}
		fmt.Printf("Changed % to %s", FilePath1, FilePath2)
		err = os.Rename(FilePath3, FilePath1)
		if err != nil {
			fmt.Printf(err.Error())
		}
		fmt.Printf("Changed % to %s", FilePath3, FilePath1)
		err = os.Rename(FilePath2, FilePath3)
		if err != nil {
			fmt.Printf(err.Error())
		}

		fmt.Printf("Changed % to %s", FilePath2, FilePath3)

		time.Sleep(100)
	}
}
