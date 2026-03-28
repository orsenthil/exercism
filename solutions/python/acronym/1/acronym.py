import re

def abbreviate(words):
    """
    Abbreviate a string.
    It removes underscores, and splits by whitespaces, plus, and hyphens. It
    returns the first letter of each word in uppercase.
    """
    words = re.sub(r'[_]', '', words)
    words = re.sub(r'[\s+-]', ' ', words)
    return ''.join(word[0].upper() for word in words.split())