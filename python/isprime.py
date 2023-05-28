#Is number prime?

#Isprime... return true only if argument number is prime
def is_prime(number):
    for i in range(2,number+1):
        if   i != number:
            if number % i ==0 :
                return False
    return True

def print_prime_number(ran):
    for i in range(2,ran+1):
        if is_prime(i):
            print(i)

print_prime_number(100)
