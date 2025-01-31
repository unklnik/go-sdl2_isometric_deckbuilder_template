package main

func main() {
	//BELOW >> WIN(1920X1080 WINDOW, 16px BASE UNIT, WINDOW TITLE, BLACK BACKGROUND, 60 FPS)
	WIN(1920, 1080, 16, "SDL Window", BLACK(), 60)
	PLAY()
	CORE()
}
func PLAY() {
	// ████ RUN ONCE CODE HERE ████ RUN ONCE CODE HERE ████ RUN ONCE CODE HERE ████ RUN ONCE CODE HERE ████ RUN ONCE CODE HERE ████

	for ONOFF {
		B4DRAW() // ►► DO NOT DELETE ►► DO NOT DELETE
		// ████ RUN GAME LOOP (REPEATING) CODE HERE ████ RUN GAME LOOP (REPEATING) CODE HERE ████ RUN GAME LOOP (REPEATING) CODE HERE ████

		//TileSheetܛDRAW(rocksTERRAIN, 0, 0, 2, 2)

		if debugOn {
			GridܛDRAW(lev)
			ObjLinesܛDRAW(terrain, MAGENTA())
			//GridBlokCentersܛDRAW(lev)
		} else {
			GridTilesMouseBlokܛDRAW(lev)
			ObjܛDRAW(terrain)
			//CardSingleܛDRAW(cards[0], 10, 10, 3, true)
			CardListܛDRAW(cards, 10, WinH-200, -5, 4, true)
			PlayersܛDRAW()

		}

		// ████ END CODE HERE ████ END CODE HERE ████ END CODE HERE ████ END CODE HERE ████
		UPDATE() // ►► DO NOT DELETE ►► DO NOT DELETE
	}
}
