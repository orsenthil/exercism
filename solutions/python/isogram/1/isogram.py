def is_isogram(string):
    string = string.replace(" ", "")
    string = string.replace('-', "")
    string = string.lower()
    for i, c in enumerate(string):
        if c in string[i+1:]:
            return False
    return True

