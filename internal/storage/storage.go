package storage

import "os"

func ReadFile(path string) (string, error) {
	file, err := os.ReadFile("./.files/" + path + ".json")
	if err != nil {
		return "", err
	}

	return string(file), nil
}

func WriteFile(path string, content string) error {
	byteContent := []byte(content)

	// TODO reavaliar as permissões xD
	return os.WriteFile("./.files/"+path+".json", byteContent, 0777)
}

func DeleteFile(path string) error {
	return os.Remove("./.files/" + path + ".json")
}
