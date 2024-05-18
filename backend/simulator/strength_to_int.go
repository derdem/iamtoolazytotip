package simulator

func StrengthToInt(strength Strength) int {
	switch strength {
	case LowStrength:
		return 1
	case MediumStrength:
		return 2
	case HighStrength:
		return 3
	default:
		panic("Country Unknown strength")
	}
}
