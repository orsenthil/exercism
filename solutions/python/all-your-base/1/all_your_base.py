def rebase(input_base, digits, output_base):
    if input_base < 2:
        raise ValueError("input base must be >= 2")
        
    if output_base < 2:
        raise ValueError("output base must be >= 2")
        
    if any(d >= input_base for d in digits):
        raise ValueError("all digits must satisfy 0 <= d < input base")
        
    if any(d < 0 for d in digits):
        raise ValueError("all digits must satisfy 0 <= d < input base")
        
    number = sum(d*input_base**i for i,d in enumerate(digits[::-1]))
    if number == 0:
        return [0]
        
    new_digits = []
    while number != 0:
        new_digits.insert(0, number%output_base)
        number = number//output_base
        
    return new_digits