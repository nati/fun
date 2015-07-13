package lifegame

//MacroCell is a square of cellular automata 2n by 2n in size.
//Each macrocell stores pointers to 5 smaller 2n-1 by 2n-1 cells.
// 0 Result
// 1 2 Future Step Cells
// 3 4
type MacroCell struct {
	child [5]*MacroCell
	size  int
}
