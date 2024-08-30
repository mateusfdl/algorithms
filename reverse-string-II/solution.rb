# rubocop:disable all
def reverse_str(s, k)
  ns = ""
  i = 0

  while i < s.length
    first_k = s[i, k].reverse

    next_k = s[i + k, k]

    ns += first_k + (next_k || "")
    i += 2 * k
  end

  ns
end

puts reverse_str('abcdefghijklmnopqrstuvxyz', 3)

