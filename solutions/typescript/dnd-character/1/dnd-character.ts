export class DnDCharacter {
 // This character has, among other things, six abilities; strength, dexterity, constitution, intelligence, wisdom and charisma.

  strength: number = DnDCharacter.generateAbilityScore();
  dexterity: number = DnDCharacter.generateAbilityScore();
  constitution: number = DnDCharacter.generateAbilityScore();
  intelligence: number = DnDCharacter.generateAbilityScore();
  wisdom: number = DnDCharacter.generateAbilityScore();
  charisma: number = DnDCharacter.generateAbilityScore();
  hitpoints: number = 10 + DnDCharacter.getModifierFor(this.constitution);

  public static generateAbilityScore(): number {
    const rolls = Array.from({ length: 4 }, () => Math.floor(Math.random() * 6) + 1);
    const lowestRoll = Math.min(...rolls);
    for (let i = 0; i < rolls.length; i++) {
        if (rolls[i] === lowestRoll) {
            rolls.splice(i, 1);
            break;
        }
    }

    var sum : number = 0;
    for (let i = 0; i < rolls.length; i++) {
        sum = sum + rolls[i];
    }

    return sum;
  }

  public static getModifierFor(abilityValue: number): number {
    return Math.floor((abilityValue - 10) / 2);
  }
}
