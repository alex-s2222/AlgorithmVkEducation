"""
Найти 2 числа сумма которых будет равняться искомому числу 
* массив отсортирован в порядке увеличения
Сложность алгоритма О(n) (так как зависит от входных параметров )
"""

from typing import List


def two_sum(arr: List[int], summ: int) -> List[int]:
    left, right = 0, len(arr) - 1

    while right != left:
        temp = arr[left] + arr[right]
        if temp == summ:
            return [arr[left], arr[right]]
        if temp < summ:
            left += 1
            continue
        right -= 1

    return None


def main():
    arr = [1, 3, 5, 6, 7, 8, 9]
    summ = 4
    print(f"Число {arr} из чисел {two_sum(arr, summ)}")
    summ = 7
    print(f"Число {arr} из чисел {two_sum(arr, summ)}")


if __name__ == "__main__":
    main()
