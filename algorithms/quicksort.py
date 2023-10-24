
def lomuto(MyList, start, end):
    pivot = MyList[end]
    i = start - 1 

    for j in range (start, end):
        if pivot >= MyList[j]:
            i += 1
            MyList[i],MyList[j] = MyList[j],MyList[i]
    
    MyList[i+1],MyList[end] = MyList[end], MyList[i+1]

    return i+1

def QuickSort(MyList, start, end):
    if end > start:
        pivot = lomuto(MyList, start, end)
        QuickSort(MyList, start, pivot-1)
        QuickSort(MyList,pivot+1, end)




def orgnized(ListDisordered):
    QuickSort(ListDisordered, 0, len(ListDisordered)-1)

ListDisordered = [2,10,3,1,9,0,100,22,4,20]
orgnized(ListDisordered)
print(ListDisordered)