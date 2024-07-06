package SystemInfo

const KB_MB_ConversionFactor float64 = 1024
const KB_GB_ConversionFactor float64 = 1024 * 1024

func kiloByteTomegaByte(val float64) float64 {
	return val / KB_MB_ConversionFactor
}

func kiloByteTogigaByte(val float64) float64 {
	return val / KB_GB_ConversionFactor
}
