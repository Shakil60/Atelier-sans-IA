package ateliersansia

import "fmt"

// PrintCenteredPyramid affiche une pyramide de hauteur n, centrée horizontalement.
func PrintCenteredPyramid(n int) {
	for i := range n {
		// Calcul du nombre d'espaces avant les étoiles
		// Plus on descend, moins il y a d'espaces
		spaces := n - i - 1

		// Affiche les espaces à gauche
		for j := 0; j < spaces; j++ {
			fmt.Print(" ")
		}

		// Affiche les étoiles (le nombre augmente à chaque ligne)
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}

		// Retour à la ligne après chaque rangée
		fmt.Println()
	}
}