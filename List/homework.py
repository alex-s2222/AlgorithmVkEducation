"""
Найти ближайшую пару элементов в массиве
Дан отсортированный массив положительных целых чисел. Необходимо найти пару соседних элементов, которые имеют наименьшую абсолютную разницу.

Первым параметром на вход программы подается размер массива.
Вторым параметром – сам массив (значения указаны через пробел).

Без учета конечных случаев
"""

def transform(elements: list) -> str:
    idx = 0
    idxx = 1
    return_list = [str(elements[idx]), str(elements[idxx])]
    min_dif = elements[idxx] - elements[idx]
    
    while idxx < len(elements):
        dif = elements[idxx] - elements[idx]
        if dif < min_dif:
            min_dif = dif
            return_list = [str(elements[idx]), str(elements[idxx])]
        idx += 1
        idxx += 1
    return ' '.join(return_list)


def main():
    size_arr = int(input())
    elements = list(map(int, input().split()))
    print(transform(elements))

if __name__ == "__main__":
    main()
