package cmd

import (
	"fmt"
	"io"
	"strings"
)

func printBorder(w io.Writer, passage string) {
	fmt.Fprintf(w, "+%s+\n", strings.Repeat("-", 100))
	fmt.Fprintf(w, "*** %s ***\n", strings.ToUpper(passage))
	fmt.Fprintf(w, "+%s+\n", strings.Repeat("-", 100))
}

func getQuranChapterNumber(chapter string) (int, error) {
	var err error
	var number int

	switch chapter {
	case "Al-Fatihah":
		number = 1
	case "Al-Baqarah":
		number = 2
	case "Al-ee-Imran":
		number = 3
	case "An-Nisa":
		number = 4
	case "Al-Ma'idah":
		number = 5
	case "Al-An'am":
		number = 6
	case "Al-A'raf":
		number = 7
	case "Al-Anfal":
		number = 8
	case "At-Tawbah":
		number = 9
	case "Yunus":
		number = 10
	case "Hud":
		number = 11
	case "Yusuf":
		number = 12
	case "Ar-Ra'd":
		number = 13
	case "Ibrahim":
		number = 14
	case "Al-Hijr":
		number = 15
	case "An-Nahl":
		number = 16
	case "Al-Isra":
		number = 17
	case "Al-Kahf":
		number = 18
	case "Maryam":
		number = 19
	case "Ta-Ha":
		number = 20
	case "Al-Anbiya":
		number = 21
	case "Al-Hajj":
		number = 22
	case "Al-Mu'minun":
		number = 23
	case "An-Nur":
		number = 24
	case "Al-Furqan":
		number = 25
	case "Ash-Shu'ara":
		number = 26
	case "An-Naml":
		number = 27
	case "Al-Qasas":
		number = 28
	case "Al-Ankabut":
		number = 29
	case "Ar-Rum":
		number = 30
	case "Luqmaan":
		number = 31
	case "As-Sajdah":
		number = 32
	case "Al-Ahzaab":
		number = 33
	case "Saba (surah)":
		number = 34
	case "Faatir":
		number = 35
	case "Ya-Sin":
		number = 36
	case "As-Saaffaat":
		number = 37
	case "Saad":
		number = 38
	case "Az-Zumar":
		number = 39
	case "Ghafir":
		number = 40
	case "Fussilat":
		number = 41
	case "Ash-Shura":
		number = 42
	case "Az-Zhukhruf":
		number = 43
	case "Ad-Dukhaan":
		number = 44
	case "Al-Jaathiyah":
		number = 45
	case "Al-Ahqaaf":
		number = 46
	case "Muhammad":
		number = 47
	case "Al-Fath":
		number = 48
	case "Al-Hujuraat":
		number = 49
	case "Qaaf":
		number = 50
	case "Adh-Dhaariyaat":
		number = 51
	case "At-Toor":
		number = 52
	case "An-Najm":
		number = 53
	case "Al-Qamar":
		number = 54
	case "Ar-Rahman":
		number = 55
	case "Al-Waqi'ah":
		number = 56
	case "Al-Hadeed":
		number = 57
	case "Al-Mujadila":
		number = 58
	case "Al-Hashr":
		number = 59
	case "Al-Mumtahanah":
		number = 60
	case "As-Saff":
		number = 61
	case "Al-Jumu'ah":
		number = 62
	case "Al-Munafiqoon":
		number = 63
	case "At-Taghabun":
		number = 64
	case "At-Talaq":
		number = 65
	case "At-Tahreem":
		number = 66
	case "Al-Mulk":
		number = 67
	case "Al-Qalam":
		number = 68
	case "Al-Haaqqa":
		number = 69
	case "Al-Ma'aarij":
		number = 70
	case "Nuh":
		number = 71
	case "Al-Jinn":
		number = 72
	case "Al-Muzzammil":
		number = 73
	case "Al-Muddaththir":
		number = 74
	case "Al-Qiyamah":
		number = 75
	case "Al-Insaan|Ad-Dahr":
		number = 76
	case "Al-Mursalaat":
		number = 77
	case "An-Naba'":
		number = 78
	case "An-Naazi'aat":
		number = 79
	case "Abasa":
		number = 80
	case "At-Takweer":
		number = 81
	case "Al-Infitar":
		number = 82
	case "At-Taffeef":
		number = 83
	case "Al-Inshiqaaq":
		number = 84
	case "Al-Burooj":
		number = 85
	case "At-Taariq":
		number = 86
	case "Al-A'la":
		number = 87
	case "Al-Ghaashiyah":
		number = 88
	case "Al-Fajr":
		number = 89
	case "Al-Balad":
		number = 90
	case "Ash-Shams":
		number = 91
	case "Al-Layl":
		number = 92
	case "Ad-Dhuha":
		number = 93
	case "Ash-Sharh":
		number = 94
	case "At-Tin":
		number = 95
	case "Al-Alaq":
		number = 96
	case "Al-Qadr":
		number = 97
	case "Al-Bayyinahh":
		number = 98
	case "Az-Zalzalah":
		number = 99
	case "Al-'Aadiyat":
		number = 100
	case "Al-Qaari'ah":
		number = 101
	case "At-Takaathur":
		number = 102
	case "Al-'Asr":
		number = 103
	case "Al-Humazah":
		number = 104
	case "Al-Feel":
		number = 105
	case "Quraish":
		number = 106
	case "Al-Maa'oon":
		number = 107
	case "Al-Kawthar":
		number = 108
	case "Al-Kaafiroon":
		number = 109
	case "An-Nasr":
		number = 110
	case "Al-Masad":
		number = 111
	case "Al-Ikhlaas":
		number = 112
	case "Al-Falaq":
		number = 113
	case "An-Naas":
		number = 114
	default:
		err = ErrInvalidPassageSpecified
	}

	return number, err
}
