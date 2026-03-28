
def translate(text):
    result = []
    for word in text.split():
        if word[0] in "aeiou" or word[:2] in ["yt", "xr"]:
            result.append(word + "ay")
            continue

        for idx in range(1, len(word)):
            if word[idx] in 'aeiou' + 'y':
                # break the consonant cluster with a vowel or 'y
                # Rule 4
                if word[idx-1] == 'q' and word[idx] == 'u':
                    # rule 3 for qu
                    idx += 1
                # rule 2
                result.append(word[idx:] + word[:idx] + 'ay')
                break

    return " ".join(result)
