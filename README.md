# Religio

A command-line client for reading The Holy Bible and Holy Quran in several languages. It has the following features:

- Fetches bible data from Bible Brain (https://www.faithcomesbyhearing.com/audio-bible-resources/bible-brain)
- Read the bible in different languages and versions
- Fetches quran data from Al Quran Cloud (https://alquran.cloud/)
- Read the quran in different languages

## Usage

### Read the bible

```go
bible -language ENG -version ESV "2 Peter 3:1-3"

// Output:
// +----------------------------------------------------------------------------------------------------+
// *** 2 PETER 3:1-3 ***
// +----------------------------------------------------------------------------------------------------+
// (1) Amados, ésta es la segunda carta que os escribo. En estas dos cartas estimulo con exhortación vuestro limpio entendimiento,
// (2) para que recordéis las palabras que antes han sido dichas por los santos profetas, y el mandamiento del Señor y Salvador declarado por
// vuestros apóstoles.
// (3) Primeramente, sabed que en los últimos días vendrán burladores con sus burlas, quienes procederán según sus bajas pasiones,
```

### Read the quran

```go
quran -language ru Al-Fatihah:1-5

// Output:
// +----------------------------------------------------------------------------------------------------+
// *** AL-FATIHAH:1-5 ***
// +----------------------------------------------------------------------------------------------------+
// (1) Во имя Аллаха, Милостивого, Милосердного!
// (2) Хвала Аллаху, Господу миров,
// (3) Милостивому, Милосердному,
// (4) Властелину Дня воздаяния!
// (5) Тебе одному мы поклоняемся и Тебя одного молим о помощи.
```
