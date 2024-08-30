def level_order(root)
  return [] if root.nil?
  result = []
  queue = [root]

  while queue.length > 0 
      level = queue.length
      currentLevel = []
      i = 0

      while i < level
          node = queue.shift
          currentLevel << node.val
          queue.push(node.left) if node.left
          queue.push(node.right) if node.right
          i+=1
      end
      result << currentLevel
  end
  result
end

# Time complexity: O(n)
# Space complexity: O(n)