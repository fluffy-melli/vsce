package package_manager

import (
	"encoding/json"
	"fmt"
	"os"
)

type MODULE struct {
	Name   string   `json:"name"`
	MainF  string   `json:"main-file"`
	Depend []string `json:"depend"`
}

func New(name string) *MODULE {
	mainF := "main.vc"
	module := &MODULE{
		Name:   name,
		MainF:  mainF,
		Depend: []string{},
	}
	jsonData, err := json.MarshalIndent(module, "", "    ")
	if err != nil {
		fmt.Println("JSON Comfile Fail:", err)
	}
	file, err := os.Create(mainF)
	if err != nil {
		fmt.Println("File Creaft Fail:", err)
		return nil
	}
	defer file.Close()
	file, err = os.Create("vsce.package")
	if err != nil {
		fmt.Println("File Creaft Fail:", err)
		return nil
	}
	defer file.Close()
	_, err = file.WriteString(string(jsonData))
	if err != nil {
		fmt.Println("File Write Fail:", err)
		return nil
	}
	return module
}

func Read() *MODULE {
	data, err := os.ReadFile("vsce.package")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return nil
	}
	var module *MODULE
	err = json.Unmarshal(data, module)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return nil
	}
	return module
}

func (m *MODULE) Download() {
	if m == nil {
		return
	}
	for _, module := range m.Depend {
		fmt.Println("Download Module : ", module)
		file, url := Domain(module)
		Download(file, url)
		os.Remove(file)
	}
}
