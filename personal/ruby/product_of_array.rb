# Given an array of numbers, replace each number with the product of all the numbers in the array except the number itself *without* using division.
#
# input = [1, 2, 3, 4, 5]
# expected  = [120, 60, 40, 30, 24]

def solution(arr)
  arr.map do |i|
    arr_without_i = arr.select { |item| item != i }
    arr_without_i.reduce { |n, acc| acc *= n }
  end
end

p solution([1, 2, 3, 4, 5])
