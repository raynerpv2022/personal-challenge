import re

def gg(str1,str2):
    return str1 == str2

def anagram(str1,str2):
    # An anagram is a word or phrase formed by rearranging the letters of a different word or phrase
    lowerstr1 = str1.lower()
    lowerstr2 = str2.lower()
    if (len(str1) != len(str2)) or (lowerstr1 == lowerstr2):
        return False
    for i in lowerstr1:
        if i not in lowerstr2:
            return False
    return True


def palindrom(str1,str2):
    # A palindrome is a word, number, phrase, or other sequence of symbols that reads the same backwards as forwards
    pattern = r"[^a-z0-9]"
    lowerstr1 = str1.lower()
    lowerstr2 = str2.lower()
    lowerstr1 = re.sub(pattern,"",lowerstr1)
    lowerstr2 = re.sub(pattern,"",lowerstr2)
    
    if len(str1) != len(str2):
        return False
    reverStr = ""
    
    for i in lowerstr1:
        reverStr = i + reverStr
        # print(reverStr)
     
    return reverStr == lowerstr2

# while True:
#     str1  = input("Enter first string ... 'q' to quit: ")
#     if str1 == "q":
#         break
#     str2  = input("Enter second string ... 'q' to quit: ")
#     if str2 == 'q':
#         break
    
#     print(str1 + "---" + str2 + ":") 
#     if palindrom(str1,str2):
#         print("Palindrom")  
#     else:
#         print("Not Palindrom")  

#     if anagram(str1, str2):
#         print("Anagram")
#     else:
#         print("Anagram")