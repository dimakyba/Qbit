def second_max_index(lst):
    if len(lst) < 2:
        return "List must have at least two elements"

    max_num = max(lst[0], lst[1])
    second_max = min(lst[0], lst[1])
    max_index = lst.index(max_num)
    second_max_index = 1 if max_num == lst[0] else 0

    for i in range(2, len(lst)):
        if lst[i] > max_num:
            second_max = max_num
            second_max_index = max_index
            max_num = lst[i]
            max_index = i
        elif lst[i] == max_num:
            max_index = i
        elif lst[i] > second_max:
            second_max = lst[i]
            second_max_index = i

    return second_max_index


lst = [10, 2, 4, 5, 0, 2, 6, 10, -1, 0, -2]
print(f'{lst[second_max_index(lst)]} {second_max_index(lst)}')
lst = [-2, 2, 0, -3]
print(f'{lst[second_max_index(lst)]} {second_max_index(lst)}')
lst = [-2, 2, 0, 0]
print(f'{lst[second_max_index(lst)]} {second_max_index(lst)}')
