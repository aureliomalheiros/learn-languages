"""
Binary search
"""

#Search value
def binary_search(list, number):
    
    start = 0
    end   = len(list) - 1
    while start <= end:
        mid = int((start + end) / 2)
        pick = list[mid]
        if pick == number:
            end = mid - 1
            return mid
        else:
            start = mid + 1
            return None
        
if __name__ == "__main__":
    my_list = [1, 2, 3, 5, 7, 9]

    pick = int(input("Pick: "))

    print(binary_search(my_list, pick))
