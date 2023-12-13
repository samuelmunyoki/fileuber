package server

import (
	"fmt"
	"math"
	"os"
)

type Directory struct{
	Dirname string
	IsDir bool
	Size string
}

var directories []Directory;


// https://stackoverflow.com/a/1094933/1333724
func PrettyByteSize(b int64) string {
	bf := float64(b)
	for _, unit := range []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi"} {
		if math.Abs(bf) < 1024.0 {
			return fmt.Sprintf("%3.1f%sB", bf, unit)
		}
		bf /= 1024.0
	}
	return fmt.Sprintf("%.1fYiB", bf)
}
func GetDirectories(scfgPath string)( []Directory, error){
	
	files, err := os.ReadDir(scfgPath)

    if err != nil {

		return  nil, err
    }

	directories= nil

    for _, file := range files {

		fi, _ := file.Info()

		newDir := Directory{

			Dirname: file.Name() ,

			IsDir: file.IsDir(),

			Size:PrettyByteSize(fi.Size())}

		directories = append(directories, newDir)	
    }
	
	return directories, nil
}
