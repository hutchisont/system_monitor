package SystemInfo

const KB_MB_ConversionFactor float64 = 1000
const KB_GB_ConversionFactor float64 = 1000 * 1000

func kiloByteTomegaByte(val float64) float64 {
	return val / KB_MB_ConversionFactor
}

func kiloByteTogigaByte(val float64) float64 {
	return val / KB_GB_ConversionFactor
}
