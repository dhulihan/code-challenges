class Node
  attr_accessor :value, :left, :right
  def initialize(value, left = nil, right = nil)
    raise "Must contain value" unless value

    @value = value
    @left = left
    @right = right
  end

  def pre_order(values = [])
    if value.nil?
      return values
    end

    values << value

    if root.left != nil
      preorder(values)
    end

    if root.right != nil
      preorder(values)
    end

    return values
  end

  def in_order(values = [])
    if value.nil?
      return values
    end

    if root.left != nil
      preorder(values)
    end

    values << value

    if root.right != nil
      preorder(values)
    end

    return values
  end

  def post_order(values = [])
    if value.nil?
      return values
    end

    if root.left != nil
      preorder(values)
    end

    if root.right != nil
      preorder(values)
    end

    values << value

    return values
  end

  # Remember, integers are not mutated when passed as an argument
  def levels(level_counter = 0)
    level_counter += 1
    #puts "value: #{value}, level_counter: #{level_counter}"

    if left.nil? && right.nil?
      return level_counter
    end

    left_levels = left ? left.levels(level_counter) : 0
    right_levels = right ? right.levels(level_counter) : 0

    if left_levels > right_levels
      return left_levels
    else
      return right_levels
    end
  end

  def search(search_value)
    #puts "search_value: #{search_value}, current_value: #{value}"
    if search_value == value
      return self
    elsif search_value < value
      left ? left.search(search_value) : nil
    elsif search_value > value
      right ? right.search(search_value) : nil
    end
  end

  def search(search_value)
    #puts "search_value: #{search_value}, current_value: #{value}"
    if search_value == value
      return self
    elsif search_value < value
      left ? left.search(search_value) : nil
    elsif search_value > value
      right ? right.search(search_value) : nil
    end
  end
end

class BinaryTree
  attr_accessor :root
  def self.generate(depth)

  end

  def initialize(root_node)
    @root = root_node
  end

  def depth
    return root.levels
  end

  def insert(node, value)
    if node.value == nil
      return node = Node.new(value)
    elsif node.value == value
      return node
    elsif value < node.value
      insert(node.left, value)
    elsif value > node.value
      insert(node.right, value)
    else
      raise "WTF!? node: #{node} value: #{value}"
    end
  end

  def search(value)
    root.search(value)
  end

  def delete(value)

  end
end

class BinarySearchTree < BinaryTree
  # check if this is a valid BST
  def valid?(node, min=-1.0/0.0, max=1.0/0.0)
    puts ("checking #{node.value}, left: #{node.left.nil?}, right: #{node.right.nil?}")
    until node.left.nil? && node.right.nil?
      puts "#{
      if min > node.value || max < node.value
        return false
      else
        valid?(node.left, min, node.value)
        valid?(node.right, node.value, max)
      end
    end
    return true
  end
end

