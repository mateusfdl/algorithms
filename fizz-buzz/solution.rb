numbers = *(1..22)

def fizz_buzz(numbers)
  i = 0
  while i < numbers.length
    result = ''
    result += 'fizz' if (numbers[i] % 3).zero?
    result += 'buzz' if (numbers[i] % 5).zero?

    pp result unless result.empty?
    pp numbers[i] if result.empty?
    i += 1

  end
end

fizz_buzz(numbers)
