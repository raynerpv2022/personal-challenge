#!/usr/bin/env python3
import anagram  
import unittest

class TestPalindrom(unittest.TestCase):
    # def test_palindrom(self):
    #     str1 = "Was it a car or a cat I saw"
    #     str2 = "Was it a car or a cat I saw"
    #     self.assertTrue(anagram.palindrom(str1,str2))
    def test_anagram(self):
        str1 =["zorra","rosa","ramo"]
        str2 = ["arroz","osar","amors"]
        for i in range(len(str1)):
            with self.subTest(i=i):
                self.assertTrue(anagram.anagram(str1[i],str2[i]))
        
    
# unittest.main()
if __name__ == "__main__":
    unittest.main()