# Добавление единицы к числу, представленному в виде массива цифр

Учитывая неотрицательное число, представленное в виде массива цифр,добавьте 1 к числу
(увеличьте число, представленное цифрами).Цифры хранятся таким образом, что самая значимая цифра является первым элементом массива.

## Пример : 
> Input : [1, 2, 4]
> 
> Output : [1, 2, 5]

## Подход: Чтобы добавить единицу к числу, представленному цифрами, выполните следующие действия :

* Разберите данный массив с конца.
* Если последний элемент 9, то он превращается 0 а перед ним добавляется 1 . `[9] => [10]`
* Для следующей итерации проверьте перенос, и если он увеличится до 10, выполните то же самое, что и шаг 2.
* После добавления переноса, сделайте перенос = 0 для следующей итерации.
* Если векторы складываются и увеличивают размер вектора, добавьте 1 в начале.
* Ниже приведена реализация добавления 1 к числу, представленному цифрами.

