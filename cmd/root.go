package cmd

import (
	"budget_planner/logger"
	"budget_planner/repository"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "budgetPlanner",
	Short: "Use budgetPlanner to gain an accurate status of your budget",
	Long:  `Use budgetPlanner to gain an accurate status of your budget`,
	Run: func(cmd *cobra.Command, args []string) {
		chooseBudget()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func chooseBudget() {
	menuOptions := []string{
		"Choose Budget",
		"New Budget",
		"Exit",
	}
	for {
		menuPrompt := promptui.Select{
			Label: "Select an option",
			Items: menuOptions,
		}
		_, result, err := menuPrompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		budgetId := -1
		// Handle the selected menu option
		switch result {
		case "Choose Budget":
			budgetId = promptAllBudget()
			whatToDoPrompt(budgetId)
		case "New Budget":
			budgetId = promptNewBudget()
			whatToDoPrompt(budgetId)
		case "Exit":
			logger.Log.Info("Exiting...")
			return
		}
	}
}

func promptNewBudget() int {
	prompt := promptui.Prompt{
		Label: "Budget Name",
	}
	name, err := prompt.Run()
	if err != nil {
		fmt.Println("Failed to read budget name:", err)
		return -1
	}

	budget := repository.Budget{Name: name}
	budgetId, err := budget.CreateBudget()
	if err != nil {
		fmt.Println("An error occurred when creating the budget", err)
		return -1
	}

	fmt.Println("Budget created successfully!")
	return int(budgetId)
}

func promptAllBudget() int {
	allBudgets, err := repository.GetAllBudget()
	if err != nil {
		return -1
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   " $ {{ .Name | cyan }}",
		Inactive: " {{ .Name | cyan }}",
		Selected: " $ {{ .Name | red | cyan }}",
	}

	menuPrompt := promptui.Select{
		Label:     "Select a Budgets",
		Items:     allBudgets,
		Templates: templates,
		Size:      4,
	}
	i, _, err := menuPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return -1
	}
	return int(allBudgets[i].ID)
}

func whatToDoPrompt(budgetId int) {
	menuOptions := []string{
		"Revenues",
		"Expenses",
		"View budget",
		"Remove budget",
		"Exit",
	}
	for {
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
		case "Revenues":
			promptRevenues(budgetId)
		case "Expenses":
			promptExpenses(budgetId)
		case "View expenses":
			// TODO: View expenses logic
			logger.Log.Info("View expenses logic goes here")
		case "Exit":
			logger.Log.Info("Exiting...")
			return
		}
	}
}
