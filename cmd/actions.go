package cmd

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

func printBorder(w io.Writer, passage string) {
	fmt.Fprintf(w, "+%s+\n", strings.Repeat("-", 100))
	fmt.Fprintf(w, "*** %s ***\n", strings.ToUpper(passage))
	fmt.Fprintf(w, "+%s+\n", strings.Repeat("-", 100))
}

func quranChapterNumber(chapter string) (int, error) {
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

func bibleBookId(book string) (string, error) {
	var err error
	var id string

	switch strings.ToLower(book) {
	case "genesis":
		id = "GEN"
	case "exodus":
		id = "EXO"
	case "leviticus":
		id = "LEV"
	case "numbers":
		id = "NUM"
	case "deuteronomy":
		id = "DEU"
	case "joshua":
		id = "JOS"
	case "judges":
		id = "JDG"
	case "ruth":
		id = "RUT"
	case "1 sammuel":
		id = "1SA"
	case "2 samuel":
		id = "2SA"
	case "1 kings":
		id = "1KI"
	case "2 kings":
		id = "2KI"
	case "1 chronicles":
		id = "1CH"
	case "2 chronicles":
		id = "2CH"
	case "ezra":
		id = "EZR"
	case "nehemiah":
		id = "NEH"
	case "esther":
		id = "EST"
	case "job":
		id = "JOB"
	case "psalms":
		id = "PSA"
	case "proverbs":
		id = "PRO"
	case "ecclesiastes":
		id = "ECC"
	case "song of solomon":
		id = "SNG"
	case "isaiah":
		id = "ISA"
	case "jeremiah":
		id = "JER"
	case "lamentations":
		id = "LAM"
	case "ezekiel":
		id = "EZK"
	case "daniel":
		id = "DAN"
	case "hosea":
		id = "HOS"
	case "joel":
		id = "JOL"
	case "amos":
		id = "AMO"
	case "obadiah":
		id = "OBA"
	case "jonah":
		id = "JON"
	case "micah":
		id = "MIC"
	case "nahum":
		id = "NAM"
	case "habakkuk":
		id = "HAB"
	case "zephaniah":
		id = "ZEP"
	case "haggai":
		id = "HAG"
	case "zechariah":
		id = "ZEC"
	case "malachi":
		id = "MAL"
	case "matthew":
		id = "MAT"
	case "mark":
		id = "MRK"
	case "luke":
		id = "LUK"
	case "john":
		id = "JHN"
	case "acts":
		id = "ACT"
	case "romans":
		id = "ROM"
	case "1 corinthians":
		id = "1CO"
	case "2 corinthians":
		id = "2CO"
	case "galatians":
		id = "GAL"
	case "ephesians":
		id = "EPH"
	case "philippians":
		id = "PHP"
	case "colossians":
		id = "COL"
	case "1 thessalonians":
		id = "1TH"
	case "2 thessalonians":
		id = "2TH"
	case "1 timothy":
		id = "1TI"
	case "2 timothy":
		id = "2TI"
	case "titus":
		id = "TIT"
	case "philemon":
		id = "PHM"
	case "hebrews":
		id = "HEB"
	case "james":
		id = "JAS"
	case "1 peter":
		id = "1PE"
	case "2 peter":
		id = "2PE"
	case "1 john":
		id = "1JN"
	case "2 john":
		id = "2JN"
	case "3 john":
		id = "3JN"
	case "jude":
		id = "JUD"
	case "revelations":
		id = "REV"
	default:
		err = ErrInvalidPassageSpecified
	}

	return id, err
}

func bibleChapter(passage string, b *bible) []string {
	if !strings.Contains(passage, ":") {
		b.chapter = passage
		return nil
	}
	chapter := strings.Split(passage, ":")
	b.chapter = chapter[0]
	return chapter
}

func bibleVerse(passage string, b *bible) error {
	if !strings.Contains(passage, "-") {
		number, err := strconv.Atoi(passage)
		if err != nil {
			return err
		}
		b.verse = append(b.verse, number)
		return nil
	}
	verses := strings.Split(passage, "-")
	for _, verse := range verses {
		number, err := strconv.Atoi(verse)
		if err != nil {
			return err
		}
		b.verse = append(b.verse, number)
	}
	return nil
}
