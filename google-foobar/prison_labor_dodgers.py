#answer = lambda l1,l2: [x for x in l1 if x not in l2]

def diff(a, b):
    b_set = set(b)
    return [elem for elem in a if elem not in b_set]

def answer(x, y):
    xy_diff = diff(x, y)

    if len(xy_diff) > 0:
        return xy_diff[0]
    else:
        yx_diff = diff(y, x)
        if len(yx_diff) > 0:
            return yx_diff[0]
        else:
            return None


#answer = lambda x,y: [item for item in x if item not in y]
#def answer(x, y):
    #x_set = set(x)
    #y_set = set(y)
    #diff = x_set - y_set
    #return diff

x = [13, 5, 6, 2, 5]
y = [5, 2, 5, 13]

print(answer(x, y))

x = [14, 27, 1, 4, 2, 50, 3, 1]
y = [2, 4, -4, 3, 1, 1, 14, 27, 50]

print(answer(x, y))


