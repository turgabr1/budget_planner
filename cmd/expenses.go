package cmd

import (
	"budget_planner/cmd/utils"
	"budget_planner/logger"
	"budget_planner/repository"
	"fmt"
	"github.com/manifoldco/promptui"
	"strconv"
)

func promptExpenses(budgetID int) {
	menuOptions := []string{
		"Add expense",
		"Edit expense",
		"Remove expense",
		"View expenses",
		"Return to previous menu",
	}
	promptContent := utils.PromptContent{
		Label: "Select an option",
	}
	for {

		_, result, _ := utils.PromptSelect(promptContent, utils.DefaultTemplates, menuOptions)

		// Handle the selected menu option
		switch result {
		case "Add expense":
			promptAddExpense(budgetID)
		case "Edit expense":
			promptEditExpense(budgetID)
		case "Remove expense":
			promptRemoveExpense(budgetID)
		case "View expenses":
			// TODO: View expenses logic
			logger.Log.Info("View expenses logic goes here")
		case "Return to previous menu":
			return
		}
	}
}

func promptRemoveExpense(budgetID int) {
	allExpenses, err := repository.GetAllExpenses(int64(budgetID))
	if err != nil {
		return
	}
	if len(allExpenses) <= 0 {
		utils.InfoPrompt("No Expenses in current Budget")
		return
	}

	promptRemove(allExpenses)
	promptContent := utils.PromptContent{
		Label: "Select an expense to remove",
	}
	i, _, _ := utils.PromptSelect(promptContent, utils.EntryTemplates, allExpenses)
	err = allExpenses[i].Delete()
	if err != nil {
		return
	}
}

func promptEditExpense(budgetID int) {
	allExpenses, err := repository.GetAllExpenses(int64(budgetID))
	if err != nil {
		return
	}
	if len(allExpenses) <= 0 {
		utils.InfoPrompt("No Expenses in current Budget")
		return
	}
	promptContent := utils.PromptContent{
		Label: "Select an expense to edit",
	}
	i, _, _ := utils.PromptSelect(promptContent, utils.EntryTemplates, allExpenses)

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
		return
	}

	resultsAmount, err := strconv.ParseFloat(results[amount], 64)
	if err != nil {
		fmt.Println("Failed to parse amount:", err)
		return
	}
	allExpenses[i].Name = name
	allExpenses[i].Amount = resultsAmount
	allExpenses[i].Occurrence = repository.Occurrence(occurrence)

	err = allExpenses[i].Update()
	if err != nil {
		return
	}

}

func promptAddExpense(budgetID int) {
	add, err := promptAdd()
	if err != nil {
		return
	}
	add.BudgetId = int64(budgetID)
	expense := repository.Expense{Entry: add}
	err = expense.Create()
	if err != nil {
		return
	}

	fmt.Println("Expense added successfully!")
	return
}
