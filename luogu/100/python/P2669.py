def calculate_coins(k):
    m = 0
    while (m * (m + 1)) // 2 <= k:
        m += 1
    m -= 1
    a = (m * (m + 1) * (2 * m + 1)) // 6
    b = k - (m * (m + 1)) // 2
    a += b * (m + 1)
    return a
k = int(input())
print( calculate_coins(k))