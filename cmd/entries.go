package cmd

import (
	"budget_planner/cmd/utils"
	"budget_planner/repository"
	"fmt"
	"github.com/manifoldco/promptui"
	"strconv"
)

func promptRemove(allEntries interface{}) int {
	promptContent := utils.PromptContent{
		Label: "Select an expense to remove",
	}
	i, _, _ := utils.PromptSelect(promptContent, utils.EntryTemplates, allEntries)

	return i
}

func promptEdit(allEntries []*repository.Entry) int {
	promptContent := utils.PromptContent{
		Label: "Select an expense to edit",
	}
	i, _, _ := utils.PromptSelect(promptContent, utils.EntryTemplates, allEntries)

	add, err := promptAdd()
	if err != nil {
		return -1
	}
	allEntries[i].Name = add.Name
	allEntries[i].Amount = add.Amount
	allEntries[i].Occurrence = add.Occurrence

	return i
}

func promptAdd() (*repository.Entry, error) {
	var name = "name"
	var amount = "amount"

	var inputs = []string{name, amount}
	results := utils.PromptGetInput(inputs)

	menuPrompt := promptui.Select{
		Label: "Select an occurrence",
		Items: repository.OccurrenceList(),
	}
	_, occurrence, err := menuPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return &repository.Entry{}, err
	}

	resultsAmount, err := strconv.ParseFloat(results[amount], 64)
	if err != nil {
		fmt.Println("Failed to parse amount:", err)
		return &repository.Entry{}, err
	}

	fmt.Println("Expense added successfully!")
	return &repository.Entry{Name: results[name], Amount: resultsAmount, Occurrence: repository.Occurrence(occurrence)}, nil
}
