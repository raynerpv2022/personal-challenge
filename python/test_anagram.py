#!/usr/bin/env python3
import anagram  
import unittest

class TestPalindrom(unittest.TestCase):
    def test_palindrom(self):
        str1 = "Was it a car or a cat I saw"
        str2 = "Was it a car or a cat I saw"
        self.assertTrue(anagram.palindrom(str1,str2))

    def test_anagram(self):
        str1 =["zorra","rosa","ramo"]
        str2 = ["arroz","osar","amor"]
        # for loop to test several case, introduccing subTest function, here i dont know  what is i=i argumnnet
        for i in range(len(str1)):
            with self.subTest(i=i):
                self.assertTrue(anagram.anagram(str1[i],str2[i]))
        
    
# i cant figure out why when i run test execute all anagram.py , i expected only the function to be testedq
if __name__ == "__main__":
    unittest.main()