def Fizz_Buzz(number):
    for i in range(number[2]):                                                                                                                                                                                   
        if (i % number[0]) == 0 and (i % number[1] == 0):
            print("FIZZBUZZ")
        elif (i % number[0] == 0):
            print("FIZZ")
        elif (i % number[1] == 0 ):
            print("BUZZ")
        else:
            print(i)
#Fizz_Buzz([3,5,100])

def cal_storage(filesize):
    block_size = 4096
    full_blocks = filesize // block_size
    print("full block",full_blocks)
    partial_block_remainder = filesize % block_size 
    print(" partial", partial_block_remainder)
    if partial_block_remainder > 0:
        return (full_blocks +1)*block_size 
    return block_size
print(cal_storage(1))
print(cal_storage(4096))
print(cal_storage(4097))
print(cal_storage(6000))

# ******************************************************

def even_number():
    for i in range(10):
            print(i+i)
even_number()

# **********************************************************
#recursion
def count(group):
    c = 0
    for menber in [99,88,77,66]:
        c+= 1
        if menber in
