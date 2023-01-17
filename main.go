package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

func main() {
	validate := func(input string) error {
		// check container name
		return nil
	}
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     "Input Container Name",
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	fmt.Printf("You choose container: %q\n", result)

	promptSelect := promptui.Select{
		Label: "Select the detect items",
		Items: []string{"CPU", "Memory", "Network", "I/O", "System"},
	}
	_, result, err = promptSelect.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	fmt.Printf("You choose: %q\n", result)
}
