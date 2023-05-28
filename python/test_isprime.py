import unittest, isprime

#  defining a class test
class TestIsprime(unittest.TestCase):
    def test_prime(self):
        number  = [0,1,2,3,4,5,6,7,8,9,10]
        expected = [False,False,True,True,False,True,False,True,False,False,False]
        for i in range(len(number)):
              with self.subTest(i=i):
                 self.assertEqual(isprime.is_prime(number[i]),expected[i])
# same problem why exectute isprime.py, i expected only test function
if __name__ == "__main__":
    unittest.main()
