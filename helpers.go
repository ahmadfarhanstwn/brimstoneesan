package brimstoneesan

import "os"

func (b *Brimstoneesan) CreateDirIfNotExist(path string) error {
	const mode = 0755
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, mode)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Brimstoneesan) CreateFileIfNotExist(path string) error {
	const mode = 0755
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return err
		}

		defer func(file *os.File) {
			file.Close()
		}(file)
	}

	return nil
}
