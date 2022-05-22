def timeConversion(s)
    arr1 = (1..12).to_a 
    arr2 = (13..23).to_a << 12
    time_hour_map_converter = arr1.zip(arr2).to_h
    
    if s.end_with?("PM") 
        s[0..1] = time_hour_map_converter[s[0..1].to_i].to_s
    end
    
    if s.start_with?("12") && s.end_with?("AM") && s[3..4].to_i >= 0
        s[0..1] = "00"
    end
    s.gsub!(/PM|AM/, "")
end

pp timeConversion("07:05:45PM")