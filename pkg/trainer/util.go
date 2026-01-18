package trainer

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const procDir = "/proc"

func findProcessByName(name string) (int, error) {
	procs, err := os.ReadDir(procDir)
	if err != nil {
		return 0, err
	}

	for _, proc := range procs {
		if !proc.IsDir() {
			continue
		}
		data, err := os.ReadFile(fmt.Sprintf("%s/%s/stat", procDir, proc.Name()))
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return 0, err
		}
		stats := strings.Split(string(data), " ")
		res := strings.Trim(stats[1], "()")
		if res == name {
			return strconv.Atoi(proc.Name())
		}
	}
	return 0, fmt.Errorf("process %s not found", name)
}
