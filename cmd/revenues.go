package cmd

import (
	"budget_planner/cmd/utils"
	"budget_planner/logger"
	"budget_planner/repository"
	"fmt"
	"github.com/manifoldco/promptui"
	"strconv"
	"strings"
)

func promptRevenues(budgetID int) {
	menuOptions := []string{
		"Add revenue",
		"Edit revenue",
		"Remove revenue",
		"View revenues",
		"Return to previous menu",
	}
	for {
		// Display the main menu and prompt the user to select an option
		menuPrompt := promptui.Select{
			Label: "Select an option",
			Items: menuOptions,
		}
		_, result, err := menuPrompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		// Handle the selected menu option
		switch result {
		case "Add revenue":
			promptAddRevenue(budgetID)
		case "Edit revenue":
			promptEditRevenue(budgetID)
		case "Remove revenue":
			promptRemoveRevenue(budgetID)
		case "View revenues":
			// TODO: View expenses logic
			logger.Log.Info("View revenues logic goes here")
		case "Return to previous menu":
			return
		}
	}
}

func promptRemoveRevenue(budgetID int) {
	allRevenues, err := repository.GetAllRevenues(int64(budgetID))
	if err != nil {
		return
	}
	if len(allRevenues) <= 0 {
		utils.InfoPrompt("No revenues in current Budget")
		return
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   " $ {{ .Name | cyan }} ({{ .Amount | red }})",
		Inactive: " {{ .Name | cyan }} ({{ .Amount | red }})",
		Selected: " $ {{ .Name | red | cyan }}",
		Details: `
--------- Revenues ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Amount:" | faint }}	{{ .Amount }}
{{ "Occurrence:" | faint }}	{{ .Occurrence }}`,
	}

	searcher := func(input string, index int) bool {
		revenue := allRevenues[index]
		name := strings.Replace(strings.ToLower(revenue.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}
	menuPrompt := promptui.Select{
		Label:     "Select an revenue to remove",
		Items:     allRevenues,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}
	i, _, err := menuPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	err = allRevenues[i].Delete()
	if err != nil {
		return
	}
}

func promptEditRevenue(budgetID int) {
	allRevenues, err := repository.GetAllRevenues(int64(budgetID))
	if err != nil {
		return
	}
	if len(allRevenues) <= 0 {
		prompt := promptui.Prompt{
			Label:     "No Revenues in current Budget",
			IsConfirm: true,
		}

		_, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		return
	}
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   " $ {{ .Name | cyan }} ({{ .Amount | red }})",
		Inactive: " {{ .Name | cyan }} ({{ .Amount | red }})",
		Selected: " $ {{ .Name | red | cyan }}",
		Details: `
--------- Revenues ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Amount:" | faint }}	{{ .Amount }}
{{ "Occurrence:" | faint }}	{{ .Occurrence }}`,
	}

	searcher := func(input string, index int) bool {
		revenue := allRevenues[index]
		name := strings.Replace(strings.ToLower(revenue.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}
	menuPrompt := promptui.Select{
		Label:     "Select an expense to remove",
		Items:     allRevenues,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}
	i, _, err := menuPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	prompt := promptui.Prompt{
		Label:   "Name",
		Default: allRevenues[i].Name,
	}
	name, err := prompt.Run()
	if err != nil {
		fmt.Println("Failed to read name:", err)
		return
	}
	allRevenues[i].Name = name

	prompt = promptui.Prompt{
		Label:   "Amount",
		Default: fmt.Sprintf("%.2f", allRevenues[i].Amount),
	}
	amountStr, err := prompt.Run()
	if err != nil {
		fmt.Println("Failed to read amount:", err)
		return
	}
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Println("Failed to parse amount:", err)
		return
	}
	allRevenues[i].Amount = amount

	promptOccurence := promptui.Select{
		Label: "Select an occurrence",

		Items: repository.OccurrenceList(),
	}
	_, occurrence, err := promptOccurence.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	allRevenues[i].Occurrence = repository.Occurrence(occurrence)

	err = allRevenues[i].Update()
	if err != nil {
		return
	}

}

func promptAddRevenue(budgetID int) string {
	prompt := promptui.Prompt{
		Label: "Name",
	}
	name, err := prompt.Run()
	if err != nil {
		fmt.Println("Failed to read name:", err)
		return ""
	}

	prompt = promptui.Prompt{
		Label: "Amount",
	}
	amountStr, err := prompt.Run()
	if err != nil {
		fmt.Println("Failed to read amount:", err)
		return ""
	}
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Println("Failed to parse amount:", err)
		return ""
	}

	menuPrompt := promptui.Select{
		Label: "Select an occurrence",
		Items: repository.OccurrenceList(),
	}
	_, occurrence, err := menuPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	expense := repository.Revenue{Entry: &repository.Entry{Name: name, Amount: amount, Occurrence: repository.Occurrence(occurrence), BudgetId: int64(budgetID)}}
	err = expense.Create()
	if err != nil {
		return ""
	}

	fmt.Println("Revenue added successfully!")
	return ""
}
