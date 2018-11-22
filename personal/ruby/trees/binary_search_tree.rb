class Node
  attr_accessor :value, :left, :right
  def initialize(value, left = nil, right = nil)
    raise "Must contain value" unless value

    @value = value
    @left = left
    @right = right
  end

  def inspect
    "<#{object_id} #{value.inspect} #{left.inspect} #{right.inspect}>"
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

  def inspect
    root.inspect
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
    node = search(value)
    remove(node) if node
  end

  def remove(node)
    puts "removing #{node.inspect}"
    if node.left.nil? && node.right.nil?
      node = nil
    # if there's only a left node, replace
    elsif !node.left.nil? && node.right.nil?
      node = node.left
    # if there's only a right node, replace
    elsif node.left.nil? && !node.right.nil?
      node = node.right
    # delete using easy method
    else
      node = delete_node_with_two_children(node)
    end
    node
  end
end

class BinarySearchTree < BinaryTree
  def delete_node_with_two_children(node)
    # find the minimum node on the right
    min_node, parent = find_min_node(node.right)
    puts("min node under #{node.inspect} is #{min_node.inspect}, parent: #{parent.inspect}")
    replace_value(node, min_node)
    remove_min_node(min_node, parent)
  end

  # We must return the parent so we can delete the original node (which
  # will now be a duplicate after we copy its value)
  def find_min_node(node, parent = nil)
    #require 'pry'
    #binding.pry
    if node.left.nil?
      return node, parent
    else
      find_min_node(node.left, node)
    end
  end

  def replace_value(old_node, new_node)
    puts("replacing #{old_node.inspect} with #{new_node.inspect}")
    old_node.value = new_node.value

    # must delete old dangling node here
  end

  def remove_min_node(node, parent)
    node = nil
    parent.left = node
  end

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
