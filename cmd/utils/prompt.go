package utils

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
)

type PromptContent struct {
	ErrorMsg string
	Label    string
	details  string
}

type InputsContent struct {
	ErrorMsg      string
	Label         string
	possibleValue string
}

var EntryTemplates = &promptui.SelectTemplates{
	Label:    "{{ . }}?",
	Active:   " $ {{ .Name | cyan }} ({{ .Amount | red }})",
	Inactive: " {{ .Name | cyan }} ({{ .Amount | red }})",
	Selected: " $ {{ .Name | red | cyan }}",
	Details: `
--------- Expense ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Amount:" | faint }}	{{ .Amount }}
{{ "Occurrence:" | faint }}	{{ .Occurrence }}`,
}

var DefaultTemplates = &promptui.SelectTemplates{
	Label:    "{{ . }}?",
	Active:   " $ {{ . | cyan }}",
	Inactive: " {{ . | cyan }}",
	Selected: " $ {{ . | red | cyan }}",
}

func PromptSelect(pc PromptContent, template *promptui.SelectTemplates, items interface{}) (int, string, error) {
	var index int
	var result string
	var err error

	menuPrompt := promptui.Select{
		Label:     pc.Label,
		Items:     items,
		Templates: template,
	}
	index, result, err = menuPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return -1, "", err
	}
	return index, result, nil
}

func PromptGetInput(inputs []string) map[string]string {
	var inputHash = map[string]string{}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	for _, input := range inputs {
		validate := func(input string) error {
			if len(input) <= 0 {
				return errors.New("Failed to validate " + input)
			}
			return nil
		}

		prompt := promptui.Prompt{
			Label:     "please enter the " + input,
			Templates: templates,
			Validate:  validate,
		}
		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			os.Exit(1)
		}
		inputHash[input] = result
	}

	return inputHash
}

func InfoPrompt(label string) {
	prompt := promptui.Prompt{
		Label: label,
	}
	_, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
}
