package fixture

func applyMasterToFade(fade int, master int) int {
	result := (float64(fade) / 100) * (float64(master) / 2.55)
	return int(result)
}
