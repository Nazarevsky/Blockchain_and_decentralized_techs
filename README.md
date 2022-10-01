# Practise_1-Big_Numbers

Обрання варіантів у додатку здійснюється за рахунок натискання потрібної кнопки (цифра, або ‘*’) та натисканні клавіші Enter

Для завдань 1, 2, 3 використовується згенерована хеш-таблиця з кількістю можливих ключів. Генерується вона один раз при обранні якогось з варіантів завдання (задля збільшення ефективності програми). 

Короткий опис функцій:
keyHandler() – функція для обробки кнопок у головному меню.
ClearConsole() – очистка консолі (працює для ОС типу windows, linux, Darwin, в іншому випадку просто відступає 3 рядка від попереднього виводу)
fillKeyCount() – заповнює хеш-таблицю з кількістю варіантів ключів по n-бітної послідовності
task1() – виконання першого завдання
task2() – виконання другого завдання
bruteForce() – шукання значення, згенерованого випадково
task3() – виконання третього завдання
