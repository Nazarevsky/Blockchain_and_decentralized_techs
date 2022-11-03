# Practise_3-Hashing_algorithms
Обрання варіантів у додатку здійснюється за рахунок натискання потрібної кнопки (цифра) та натисканні клавіші Enter</br>
Роботу поділено на три файли: main.go та SHA1.go (/hashing/SHA1.go) Keccak.go (/hashing/Keccak.go), де SHA1.go - бібліотека для хешування значень за алгоритмом SHA1,  Keccak.go - бібліотека для хешування значень за алгоритмом Keccak (наразі непрацює. В алгоритмі начебто розібрався, написав, але вже вичерпав весь час, що відводив собі на це завдання, тому не виправив помилки), main.go - CLI для використання бібліотек<br/>
<br/>
Ввод значень<br/>
На вход задається повідомлення (однією строкою) будь-якої довжинию
<br/>
Вивід значень<br/>
На вивід видається унікальне повідомлення фіксованої довжини
<br/>
Короткий опис функцій файлу SHA1.go<br/>
```binary(s string) string``` - перевод повідомлення у двійковий формат<br/>
```addZerosTo(mes string, to int) string``` - додавання певної кількості нулів до двійкового значення<br/>
```filler(binMes string) string``` - доповнення бітового значення нулями до довжини, кратної 512<br/>
```rotate(val uint32, k int) uint32``` - циклічний здвиг бітового рядка на k бітів<br/>
```SHA1(message string) string``` - функція гешування за алгоритмом SHA1<br/>
<br/>
Короткий опис функцій файлу Keccak.go<br/>
```binaryK(s string) string``` - перевод повідомлення у двійковий формат (аналогічно до алгоритму SHA1, але продубльовано, щоб не було залежностей між файлами)<br/>
```padding(binMes string) string``` - доповнення бітового значення до певної кратності<br/>
```mod(num int, m int) int``` - остача від ділення (спеціально для num < 0)<br/>
```theta(a [][][]byte) [][][]byte``` - θ перестановка<br/>
```rho(a [][][]byte) [][][]byte``` - ρ перестановка<br/>
```pi(a [][][]byte) [][][]byte``` - π перестановка<br/>
```chi(a [][][]byte) [][][]byte``` - χ перестановка<br/>
```replAtInd(in string, r byte, i int) string``` - заміна значення на певному індексу в строкі<br/>
```rc(t int) byte``` - розрахунок значень констант RC<br/>
```iota_(a [][][]byte, ir int) [][][]byte``` - ι перестановка<br/>
```squeeze(a [][][]byte) string``` - "вичавлення" значень з губки<br/>
```binToHex(s string) string``` - перевод двійкового значення у шіснадцяткове<br/>
```squeezeToHex(s string) string``` - побайтовий перевод двійкового "вичавленого" з губки значення у шіснадцяткове<br/>
```Keccak(messgae string) string``` - функція гешування за алгоритмом Keccak<br/>
<br/>
Короткий опис функцій файлу main.go<br/>
```ClearConsole()``` - очистка консолі (працює для ОС типу windows, linux, Darwin, в іншому випадку просто відступає 3 рядка від попереднього виводу)<br/>
```task(fn func(string) string, algName string)``` - функція для виконання завдань 1 та 2, залежно від переданної функції<br/>
```keyHandler(key string)``` - обробник натискань на кнопки<br/>
```main()``` - точка входу<br/>