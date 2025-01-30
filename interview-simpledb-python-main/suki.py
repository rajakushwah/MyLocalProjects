def calculate_sqrt_number(number:int):
    if number <0:
        return -1
    low = 0
    high = number
    mid = 0

    while low <= high:
        mid = (low+high) // 2
        sqr= mid * mid
        if sqr == number:
            return  mid
        elif sqr<number:
            low  = mid +1
        else:
            high = mid -1

    return high
    print(high)

x= calculate_sqrt_number(4)
print(x)