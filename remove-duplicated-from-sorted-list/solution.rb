def delete_duplicates(head)
  curr = head

  while curr
    if curr.next && curr.val == curr.next.val
      curr.next = curr.next.next
    else
      curr = curr.next
    end
  end

  head
end

# Time complexity: O(n)
# Space complexity: O(1)
