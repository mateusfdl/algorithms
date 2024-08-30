# rubocop:disable all
def can_construct(target, word_dict)
  dp = Array.new(target.length + 1, false)
  dp[0] = true

  (0..target.length).each do |i|
    if dp[i]
      word_dict.each do |word|
        if target.length - i >= word.length && target[i, word.length] == word
          dp[i + word.length] = true
        end
      end
    end
  end

  dp[target.length]
end

target = "applepenapplez"
word_dict = ["apple", "pen", "z"]

can_construct(target, word_dict)
