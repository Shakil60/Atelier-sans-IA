package ateliersansia

// MaxSubArray renvoie la somme maximale d’un sous-tableau contigu
// en utilisant l’algorithme de Kadane (optimisé en O(n)).
func MaxSubArray(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    // sommeAct : somme maximale se terminant à l’indice courant
    // sommeMax : meilleure somme rencontrée jusqu’à présent
    sommeAct := nums[0]
    sommeMax := nums[0]

    // Parcours linéaire du tableau
    for i := 1; i < len(nums); i++ {
        // Si ajouter nums[i] améliore la somme courante, on continue,
        // sinon on repart à zéro (on redémarre à nums[i]).
        if sommeAct+nums[i] > nums[i] {
            sommeAct = sommeAct + nums[i]
        } else {
            sommeAct = nums[i]
        }

        // Mise à jour du maximum global
        if sommeAct > sommeMax {
            sommeMax = sommeAct
        }
    }

    return sommeMax
}
