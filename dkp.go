package everquest

// DKPuser is an interface for returning, giving, and removing dkp
type DKPuser interface {
	getDKP() int
	giveDKP(int) int
	removeDKP(int) int
}
