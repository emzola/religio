# religio

religio is an command-line client for reading the bible and quran in several languages.

## Usage

### Read the bible

```go
bible -lang es "2 Peter 3:1-5"

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
quran -lang ru Al-Fatihah:1-5

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
