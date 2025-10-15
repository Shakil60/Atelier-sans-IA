package ateliersansia

import "fmt"

// PrintHollowDiamond affiche un losange creux composé d’astérisques (*),
// centré et symétrique. Le rayon r détermine sa taille verticale (2*r + 1 lignes).
func PrintHollowDiamond(r int) {
	// Partie supérieure (pyramide creuse)
	for i := 0; i <= r; i++ {
		// Espaces avant les bords du losange
		for j := 0; j < r-i; j++ {
			fmt.Print(" ")
		}

		// Affichage des étoiles et espaces intérieurs
		for j := 0; j < 2*i+1; j++ {
			// On affiche une étoile uniquement sur les bords
			if j == 0 || j == 2*i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}

	// Partie inférieure (pyramide inversée creuse)
	for i := r - 1; i >= 0; i-- {
		// Espaces avant les bords
		for j := 0; j < r-i; j++ {
			fmt.Print(" ")
		}

		// Affichage des étoiles et espaces intérieurs
		for j := 0; j < 2*i+1; j++ {
			if j == 0 || j == 2*i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
}
