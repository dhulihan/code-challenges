# Problem: Assume you have a method isSubstring to check if one word is a substring of another. Given two strings s1 and s2, check if s2 is a rotation of s1 with only one call to isSubstring.
#
# Inputs:
#   ("water", "erwat") # true
def rotation?(s1, s2)
  rotation = false
# strings must be of equal lengths
  if s1.length == s2.length
    # count the number of rotations
    count = 0

    # rotate for only the length of the string
    until count == s1.length
      # create array of characters
      temp = s1.split('')

      # reconcatenate chars after pushing the first char (.shift) to end (.push)
      s1 = temp.push(temp.shift).join

      if s1 == s2
        # set rotation flag to true
        rotation = true
      end

      count += 1
    end
  else
    false
  end
  rotation
end

# fun extra method that will rotate a string one position to the left
# rotate("water") # aterw
def rotate(str)
  temp = str1.chars
  temp.push(temp.shift).join
end

p rotation?("water", "erwat")
