# Test_task
Задача 1

Пусть дана некоторая функция GetResult(x int64) (int64, error) которая может выполняться достаточно долгое время, и кроме того, может задействовать большое число ресурсов
Клиент (например, фронтенд) отправляет набор целых чисел,  для которых нужно посчитать результат функцией GetResult и вывести ответ в удобном клиенту формате. Учесть, что для адекватной работы клиент готов ждать не более N секунд , в противном случае он хочет получить какую-либо информацию о том, что для некоторых из отправлнных чисел подсчет незавершен и находится в процессе.  (так же для простоты можно считать, что количество различных чисел на входе сравнительно невелико, меньше 10000 вариантов).
Задача написать программу, которая по запрошенному набору чисел выдаст результат для каждого из них. Структура ответа - на Ваше усмотрение.
Пример функции GetResult (но можно придумать свой) :
1) Ожидание случайное (или зависящее от входа) время
2) Выход = вход, ошибка пустая

Задачи 2,3

В некоторой БД (postgres) есть таблица products (id pkey, name, mark) и таблица категорий (id pkey, name)
Реализовать в этой БД отношение многие-ко-многим ,т.е. чтобы каждый продукт мог попасть в несколько категорий, и каждая категория могла включать в себя много продуктов (можно создавать новые таблицы и изменять текущие) 
Написать функции (go,sql) , первая из которых на вход принимает id продукта, а на выходе выдает всю информацию о нем (название name, оценка mark, список категорий, куда он попадает).
 Вторая на входе принимает значение «Max» или «Min» перечисляемого типа (допустимо сделать просто строкой) ,а  на выходе должна для каждой категории выдать продукт либо с максимальным рейтингом (mark), либо с минимальным


Задача 4*

Пусть даны N уникальных пользователей. Пусть они случайно разбиваются на k групп (k<<N), так, что для каждого пользователя вероятность попасть в каждую группу одинакова. Каким будет распределение случайных величин ("Количество пользователей в группе")? Можно ли его считать нормальным? Какими будут их мат ожидание и дисперсия? (Можно для примера решить для N=1000000, k=256)? Чему равна веротяность того, что число пользователей  в определенной группе  будет от 3000 до 3200?
