package helper

func CalCulateBonusStandart(jumlahPeliharaan int, qtyDeplesi int) bool {
	dataStandart := jumlahPeliharaan + (2 * jumlahPeliharaan / 100)

	return qtyDeplesi < dataStandart
}
