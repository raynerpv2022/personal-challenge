#!/usr/bin/env python3
#Is number prime?

#Isprime... return true only if argument number is prime
def is_prime(number):
    if number < 2:
        return False
    for i in range(2,number+1):
        if   i != number:
            if number % i == 0 :
                return False
    return True

# define a range to check prime number
def print_prime_number(ran):
    for i in range(2,ran+1):
        if is_prime(i):
            print(i)

print_prime_number(100)
