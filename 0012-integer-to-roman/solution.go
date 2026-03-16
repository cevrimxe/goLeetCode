func intToRoman(num int) string {

	type RomanEntry struct {
        Symbol string
        Value  int
    }

    romanTable := []RomanEntry{
        {"M", 1000},
        {"CM", 900}, // 900 özel durumu
        {"D", 500},
        {"CD", 400}, // 400 özel durumu
        {"C", 100},
        {"XC", 90}, // 90 özel durumu
        {"L", 50},
        {"XL", 40},  // 40 özel durumu
        {"X", 10},
        {"IX", 9},   // 9 özel durumu
        {"V", 5},
        {"IV", 4},   // 4 özel durumu
        {"I", 1},
    }

    var result strings.Builder

    for _, entry := range romanTable {
        for num >= entry.Value {
            result.WriteString(entry.Symbol)
            num -= entry.Value
        }
    }
    return result.String()
}
