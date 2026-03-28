export function commands(command: number): string[] {
    // convert command to binary
    //
    const binary = command.toString(2);
    const commands = [];
    if (binary.length >= 1 && binary[binary.length - 1] === '1') {
        commands.push('wink');
    }
    if (binary.length >= 2 && binary[binary.length - 2] === '1') {
        commands.push('double blink');
    }
    if (binary.length >= 3 && binary[binary.length - 3] === '1') {
        commands.push('close your eyes');
    }
    if (binary.length >= 4 && binary[binary.length - 4] === '1') {
        commands.push('jump');
    }
    if (binary.length >= 5 && binary[binary.length - 5] === '1') {
        commands.reverse();
    }
    return commands;
}
