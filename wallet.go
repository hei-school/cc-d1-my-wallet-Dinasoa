package main

import (
	"fmt"
	"strconv"
)

var transactions []string
var solde float64
var monthlyBudget float64 = 200000
var defaultUsername string = "Dinasoa"
var defaultPassword string = "averystrongpassword"

func showBalance() {
	fmt.Printf("Votre solde actuel: %.2f Ar\n", solde)
}

func deposit() {
	var amountStr string
	fmt.Print("Veuillez saisir le montant à déposer: Ar ")
	fmt.Scanln(&amountStr)

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err == nil && amount > 0 {
		solde += amount
		transactions = append(transactions, fmt.Sprintf("Dépôt de %.2f Ar", amount))
		fmt.Println("Dépôt effectué!")
	} else {
		fmt.Println("Le montant que vous avez saisi est invalide, veuillez vérifier!")
	}
}

func withdraw() {
	var amountStr string
	fmt.Print("Veuillez saisir le montant à retirer: Ar ")
	fmt.Scanln(&amountStr)

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err == nil && amount > 0 {
		if amount <= solde {
			solde -= amount
			transactions = append(transactions, fmt.Sprintf("Retrait de %.2f Ar", amount))
			fmt.Println("Retrait effectué.")
		} else {
			fmt.Println("Votre solde est insuffisant.")
		}
	} else {
		fmt.Println("Montant invalide.")
	}
}

func showTransactions() {
	fmt.Println("Historique de transactions:")
	for _, transaction := range transactions {
		fmt.Println(transaction)
	}
}

func setMonthlyBudget() {
	var budgetStr string
	fmt.Print("Veuillez saisir votre nouveau budget mensuel: Ar ")
	fmt.Scanln(&budgetStr)

	budget, err := strconv.ParseFloat(budgetStr, 64)
	if err == nil && budget > 0 {
		monthlyBudget = budget
		fmt.Printf("Votre budget mensuel a été mis à jour: %.2f Ar\n", monthlyBudget)
	} else {
		fmt.Println("Montant invalide.")
	}
}

func main() {
	var username string
	var userPassword string

	fmt.Print("Veuillez entrer votre nom d'utilisateur: ")
	fmt.Scanln(&username)
	fmt.Print("Veuillez confirmer votre identité en tapant le mot de passe: ")
	fmt.Scanln(&userPassword)

	if username != defaultUsername || userPassword != defaultPassword {
		fmt.Println("Oups, vous n'avez pas accès à ce wallet, désolé.")
	} else {
		var choice int

		for choice != 6 {
			fmt.Printf("Bienvenue dans 'Wallet' %s, veuillez choisir une action à effectuer:\n"+
				"1- Voir le solde dans mon compte\n"+
				"2- Faire un dépôt\n"+
				"3- Faire un retrait\n"+
				"4- Afficher l'historique des transactions\n"+
				"5- Mettre à jour le budget mensuel\n"+
				"6- Quitter\n", defaultUsername)

			fmt.Print("Votre choix: ")
			fmt.Scanln(&choice)

			switch choice {
			case 1:
				showBalance()
			case 2:
				deposit()
			case 3:
				withdraw()
			case 4:
				showTransactions()
			case 5:
				setMonthlyBudget()
			case 6:
				fmt.Println("Merci d'utiliser le service Wallet. Au revoir!")
			default:
				fmt.Println("Choix invalide. Veuillez choisir une option valide.")
			}
		}
	}
}
