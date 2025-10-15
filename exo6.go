package ateliersansia

type Package struct {
    Number    int
    Recipient string
    Address   string
    Content   string
    Weight    float64
}

type Wagon struct {
    Number    string
    Loading   []Package
    NextWagon *Wagon
}

type ResultSearch struct {
    NumberWagon     string
    PositionPackage int
    PackageFind     Package
}

func FindPackage(number int, startWagon *Wagon) ResultSearch {
	return
}