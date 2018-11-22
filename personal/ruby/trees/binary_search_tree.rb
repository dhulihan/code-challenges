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
  def valid?(node)
    return invalid_nodes(node).size == 0
  end

  # perform post-order validity check to ensure this is a BST tree.
  # Returns array of "invalid" nodes.
  def invalid_nodes(node, bad_nodes = [], min=-1.0/0.0, max=1.0/0.0)
    # base case
    return true if node.nil?

    #puts ("checking #{node.value}, left: #{node.left.inspect}, right: #{node.right.inspect}, min: #{min}, max: #{max}, bad_nodes: #{bad_nodes}")

    if node.value < min || node.value > max
      bad_nodes << node.value
      return false
    end

    invalid_nodes_left = invalid_nodes(node.left, bad_nodes, min, node.value)

    # uncomment this if you want to fail fast
    # return false unless left_valid

    invalid_nodes_right = invalid_nodes(node.right, bad_nodes, node.value, max)

    return bad_nodes
  end
end
