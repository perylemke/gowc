package cmd

import "os"

func CountBytes(fileName string) (int, string, error) {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		return 0, "", err
	}

	return int(fileInfo.Size()), fileInfo.Name(), nil
}
