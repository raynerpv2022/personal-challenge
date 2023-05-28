import unittest, isprime

#  defining a class test
class TestIsprime(unittest.TestCase):
    def test_prime(self):
        number  = [3,5,7]
        for i in range(len(number)):
              with self.subTest(i=i):
                 self.assertTrue(isprime.is_prime(number[i]))
# same problem why exectute isprime.py, i expected only test function
if __name__ == "__main__":
    unittest.main()
