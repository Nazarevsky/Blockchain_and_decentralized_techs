# Practise_4-Crypto_algorithms
Обрання варіантів у додатку здійснюється за рахунок натискання потрібної кнопки (цифра) та натисканні клавіші Enter</br>
Роботу поділено на п'ять файлів: main.go, AES.go (cription/AES.go), RSA.go (cription/RSA.go), Vigenere.go (cription/Vigenere.go) та arrays.go (cription/arrays.go), де AES.go - бібліотека для 
шифрування значень за алгоритмом AES, RSA - бібліотека для шифрування значень за алгоритмом RSA (шифрує тільки невеликі беззнакові числа), Vigenere - бібліотека 
для шифрування значень за алгоритмом Віженера, main.go - CLI для використання бібліотек, arrays.go - зберігання масивів для алгоритму AES.<br/>
<br/>
**Ввод значень**<br/>
Шифрування:
  - AES: значення та ключ (рядки будь-якої довжини)
  - RSA: невелике беззнакове число (проблема з переповненням)
  - Vigenere: значення та ключ (рядки будь-якої довжини, але ключ буде обрізано під довжину строки)<br/>

Розшифрування:
  - AES: шифротекст та ключ (рядки будь-якої довжини)
  - RSA: шифротекст (число)
  - Vigenere: шифротекст та ключ (рядки будь-якої довжини, але ключ буде обрізано під довжину строки)<br/>
 
**Вивід значень**<br/>
На вивід видається зашифроване значення<br/>
## Короткий опис функціоналу AES
```divMesIntoBlocks(blocks [][][]byte, bitMes string, countBlocks int) [][][]byte ``` - поділяє вхідне повідомлення на 128-ми бітні блоки<br/>

```divMesIntoBlocks(blocks [][][]byte, bitMes string, countBlocks int) [][][]byte ``` - поділяє вхідне повідомлення на 128-ми бітні блоки<br/>
```divMesIntoBlocks(blocks [][][]byte, bitMes string, countBlocks int) [][][]byte ``` - поділяє вхідне повідомлення на 128-ми бітні блоки<br/>
```intitKey(stateKey [][]byte, bitMes string) [][]byte ``` - поділяє ключ на 128-ми бітний блок<br/>
```complete(mes string, to int) string``` - додає нулі до повідомлення заданої довжини<br/>
```pad(mes string, to int) string``` - додає нулі до повідомлення, щоб його довжина була кратна заданому числу<br/>
```mesToBits(mes string) string``` - переводить рядок у бітову форму<br/>
```addKey(block [][]byte, key [][]byte) [][]byte``` - додає ключ до заданого блоку (операція XOR)<br/>
```getValInBox(hex byte, box [][]byte) byte``` - отримує значення у substitution box або inverse substitution box (залежить від переданого параметру) за 16-тковим числом<br/>
```subBytes(block [][]byte, box [][]byte) [][]byte``` - переводить значення блоку у значення, що отримано у substitution box (або inverse) <br/>
```shiftArr(arr []byte, shift int) []byte``` - здвигає значення масиву на задану кількість байт<br/>
```shiftBlock(block [][]byte) [][]byte``` - порядково здвигає значення блоку<br/>
```multHex(hex byte, mul byte) byte``` - помножує шіснадцяткові числа (на 0x01, 0x02, 0x03, 0x09, 0x0x, 0x0d, 0x0e)<br/>
```mixColumns(block [][]byte, mulmatr [][]byte) [][]byte``` - помножує заданий блок на відповідну матрицю<br/>
```genRoundKey(key [][]byte, round int) [][]byte``` - генерує раундовий ключ<br/>
```reassemble(block [][]byte) string``` - перетворює блок на рядок<br/>
```crToBlocks(blocks [][][]byte, hexMes string, countBlocks int) [][][]byte``` - поділяє зашифроване повідомлення на блоки<br/>
```blockNewInstance(block [][]byte) [][]byte``` - повертає екземпляр блоку (deep copy)<br/>
```invShiftBlock(block [][]byte) [][]byte``` - інверсивний здвиг рядків блоку (див. ```shiftBlock```)<br/>
```hexToString(block [][]byte) string``` - перетворює блок на рядок 16-ткових чисел<br/>
```AES_encrypt(mes string, key string) string``` - шифрує повідомлення за алгоритмом AES<br/>
```AES_decrypt(mes string, key string) string``` - розшифровує шифротекст за алгоритмом AES<br/>
``````
