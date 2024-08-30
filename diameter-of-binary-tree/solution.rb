class TreeNode
  attr_accessor :val, :left, :right

  def initialize(val = 0, left = nil, right = nil)
    @val = val
    @left = left
    @right = right
  end
end

root5 = TreeNode.new(1)
root5.left = TreeNode.new(2)
root5.right = TreeNode.new(3)
root5.left.left = TreeNode.new(4)
root5.left.right = TreeNode.new(5)
root5.left.right.left = TreeNode.new(6)
root5.left.right.right = TreeNode.new(7)

def diameter_of_binary_tree(root)
  _, diameter = calculate_diameter(root)
  diameter
end

def calculate_diameter(node)
  return 0, 0 unless node

  left_height, left_diameter = calculate_diameter(node.left)
  right_height, right_diameter = calculate_diameter(node.right)

  current_height = [left_height, right_height].max + 1
  current_diameter = [left_diameter, right_diameter, left_height + right_height].max

  [current_height, current_diameter]
end
