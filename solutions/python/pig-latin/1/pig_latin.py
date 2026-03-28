def is_vowel_sound(text):
    return text[0] in 'aeiou' or text.startswith('yt') or text.startswith('xr')

def is_consonant_sound(text):
    return not is_vowel_sound(text)

def is_consonant_sound_with_qu(text):
    if is_consonant_sound(text[:1]) and text[1:2] == 'qu':
        return True

def translate(text):
    if len(text) == 1:
        return text + 'ay'

    if len(text) == 2 and text[1] == 'y':
        return text + 'ay'

    if is_vowel_sound(text):
        return text + 'ay'

    if is_consonant_sound(text[:2]):
        return text[2:] + text[:2] + 'ay'

    if is_consonant_sound(text[:1]):
        return text[1:] + text[:1] + 'ay'

    if is_consonant_sound_with_qu(text):
        return text[3:] + text[:3] + 'ay'

    for i, c in enumerate(text):
        if is_consonant_sound(text[i]) and text[i+1] == 'y':
            return text[i+1:] + text[:i+1] + 'ay'

    return text
