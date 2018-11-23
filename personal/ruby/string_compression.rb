# Implement a method to perform basic string compression using the counts of repeated characters. For example, the string aabcccccaaa would become a2b1c5a3. If the “compressed” string would not become smaller than the original string, your method should return the original string. Assume the string only has uppercase and lowercase letters (a-z).

# Solution 1: use hash table
def compression(string)
  array = string.split(//)
  hash = Hash.new(0)
    array.each do |char|
     hash.include? char ? hash[char]+=1 : hash[char] = 1
  end

  str = ""
  hash.each do |k, v|
     str << "#{k}" + "#{v}"
  end
  str
end

p compression("aabbcccccdd")
p compression("ABBBaaaCCCcDdaa")

