export const answer = (question: string): number => {
    // Reject a problem other than "What is"

    if (!question.startsWith('What is')) {
        throw new Error('Unknown operation');
    }

    const matches = question.match(/What is (.*)\?/);
    console.log(matches);
    if (matches === null) {
        throw new Error('Syntax error');
    }

    const parts = matches[1].split(' ');
    console.log(parts);
    let result = parseInt(parts[0]);
    if (isNaN(result)) {
        throw new Error('Syntax error');
    }
    for (let i = 1; i < parts.length; i += 2) {
        const op = parts[i];
        if (op === 'plus') {
            const num = parseInt(parts[i + 1]);
            result += num;
        } else if (op === 'minus') {
            const num = parseInt(parts[i + 1]);
            result -= num;
        } else if (op === 'multiplied') {
            // i + 2 is the number after 'multiplied by'
            const num = parseInt(parts[i + 2]);
            i += 1;
            result *= num;
        } else if (op === 'divided') {
            // i + 2 is the number after 'divided by'
            const num = parseInt(parts[i + 2]);
            i += 1;
            result /= num;
        } else {
            // Reject a problem missing an operand
            //
            console.log(parts);
            if (op == 'cubed') {
                throw new Error('Unknown operation');
            }

            if (i + 1 >= parts.length) {
                throw new Error('Syntax error');
            }

            if (parseInt(parts[i + 1])) {
                throw new Error('Syntax error');
            }
            // Reject two operations in a row
            if (isNaN(parseInt(parts[i + 1]))) {
                throw new Error('Syntax error');
            }
            // Reject two numbers in a row
            if (isNaN(parseInt(parts[i])) || isNaN(parseInt(parts[i + 1]))) {
                throw new Error('Syntax error');
            }
        }
    }

    if (isNaN(parseInt(parts[parts.length -1]))) {
        throw new Error('Syntax error');
    }


    return result;
}
