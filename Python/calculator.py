def do_the_math(f,o,s):
    match (o):
        case "*":
            output = float(f) * float(s)
        case "/":
            output = float(f) / float(s)
        case "+":
            output = float(f) + float(s)
        case "-":
            output = float(f) - float(s)
    return output

if __name__ == "__main__":
    first_number  =input('Enter first number:')
    operation  =input('Enter operation:(+, -, *, /):')
    second_number  =input('Enter second number:')
    output = do_the_math(first_number, operation, second_number)
    print("Result:", output)