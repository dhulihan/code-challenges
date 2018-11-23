# permutations of palindrome
def palindrome(string)
  array = string.split(//)
  hash = Hash.new(0)

  array.each do |char|
    hash.include? char ? hash[char]+=1 : hash[char] = 1
  end

  odd = 0

  hash.each_value do |v|
    odd += 1 if v % 2 != 0
  end

  odd > 1 ? false : true
end

p palindrome("never odd or even")
p palindrome("racecar")
p palindrome("carrace")
p palindrome("321X321")

