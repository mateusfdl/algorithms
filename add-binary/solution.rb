def add_binary(binary1, binary2)
  binary1, binary2 = binary2, binary1 if binary1.length < binary2.length
   result = ""
   carry = 0
   (0...binary1.length).each do |i|
     digit1 = binary1[-(i + 1)].to_i 
     digit2 = binary2[-(i + 1)].to_i rescue 0 
 
     sum = digit1 + digit2 + carry
 
     carry = sum / 2
     result = (sum % 2).to_s + result
   end
 
   result = carry.to_s + result if carry > 0
 
   result
end