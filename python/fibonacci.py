# Factorial from number
def factorial_number(number):
    if number == 0:
        return 1
    return number*factorial_number(number-1)

print(factorial_number(5))

# fibonacci for loop
def fibo_for(number):
    fibo = [0,1]
    
    for i  in range(2,number+1):
        fibo.append(fibo[i-1]+fibo[i-2])
    return fibo

print(fibo_for(10))

#fibo recursive function
def fibo_recur(number):
    if number <= 1:
        return number
    return fibo_recur(number-1)+fibo_recur(number-2)

print(fibo_recur(10))