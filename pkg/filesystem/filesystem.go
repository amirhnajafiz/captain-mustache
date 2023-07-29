package filesystem

import "os"

func ReadFile(path string) (string, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return "", ErrFileNotFound
	}

	return string(f), nil
}

func WriteFile(content, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return ErrCreateFile
	}

	_, err = f.WriteString(content)

	return err
}
