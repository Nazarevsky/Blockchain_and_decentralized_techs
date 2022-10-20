Обрання варіантів у додатку здійснюється за рахунок натискання потрібної кнопки (цифра) та натисканні клавіші Enter</br>
Роботу поділено на два файли: main.go та endian.go (/endian/endian.go), де endian.go - бібліотека для перевода значень, main.go - CLI для використання бібліотеки<br/>
<br/>
Ввод значень<br/>
Прикклад вводу HEX значень: 0xfff, 0XfFf, fFfa100<br/>
При тому, значення 0xaa == 0x000...00aa<br/>
Значення Little Endian та Big Endian потрібно вводити, як звичайні десяткові беззнакові числа<br/>
<br/>
Вивід значень<br/>
HEX значення виводяться без 0x (це зроблено для зручності, бо більшість конвертерів не сприймають HEX значення з 0x на початку)<br/>
<br/>
Короткий опис функцій endian.go<br/>
```addZeroToLen(string, int, bool) string``` - додає нулів до строки, щоб її довжина була такою, яка вказана у аргументі int (додається ззаду чи спереді залежить від флагу bool)<br/>
```reverseHexNum(string) string``` - перевертає строку задом наперед по 2 значення<br/>
```HexToLittleEndian(string) *big.Int``` - переводить значення HEX у Little Endian<br/>
```HexToBigEndian(string) *big.Int``` - переводить значення HEX у Big Endian<br/>
```bigIntToHex(*big.Int) string``` - <br/>
```LittleEndianToHex(*big.Int) string``` - переводить значення Little Endian у HEX<br/>
```BigEndianToHex(*big.Int) string``` - переводить значення Big Endian у HEX<br/>
<br/>
Короткий опис функцій main.go<br/>
```ClearConsole()``` - очистка консолі (працює для ОС типу windows, linux, Darwin, в іншому випадку просто відступає 3 рядка від попереднього виводу)<br/>
```keyHandler(string)``` - функція для обробки кнопок<br/>
```isHexNum(string) (string, bool)``` - перевірка, що введене число має HEX формат<br/>
```taskFunction12(func(string) *big.Int, string)``` - функція, що призначена для виконання першого та другого завдання, залежно від переданої їй функції<br/>
```taskFunction34(func(*big.Int) string, string)``` - функція, що призначена для виконання другого та третього завдання, залежно від переданої їй функції<br/>
```main()``` - точка входу<br/>
