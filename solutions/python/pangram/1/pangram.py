def is_pangram(sentence):
    if len(sentence) < 26:
        return False
    sentence = sentence.lower()
    for letter in 'abcdefghijklmnopqrstuvwxyz':
        if letter not in sentence:
            return False
    return True