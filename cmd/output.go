package cmd

import "os"

func OutputToFile(path, data string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Write([]byte(data))
	return nil
}