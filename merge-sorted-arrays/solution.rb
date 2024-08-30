nums1 = [1, 2, 3, 0, 0, 0]
m = 3
nums2 = [2, 5, 6]
n = 3

def merge(nums1, m, nums2, n)
  ref = m + n - 1
  while ref >= m
    n -= 1
    nums1[ref] = nums2[n]
    ref -= 1
  end

  nums1.sort!
end

pp merge(nums1, m, nums2, n)
