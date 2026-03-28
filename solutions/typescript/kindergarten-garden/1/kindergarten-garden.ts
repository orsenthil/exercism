//
// This is only a SKELETON file for the 'Kindergarten Garden' exercise.
// It's been provided as a convenience to get you started writing code faster.
//

const DEFAULT_STUDENTS: Student[] = [
  'Alice',
  'Bob',
  'Charlie',
  'David',
  'Eve',
  'Fred',
  'Ginny',
  'Harriet',
  'Ileana',
  'Joseph',
  'Kincaid',
  'Larry',
]

const PLANT_CODES = {
  G: 'grass',
  V: 'violets',
  R: 'radishes',
  C: 'clover',
} as const

type Student = string
type Plant = (typeof PLANT_CODES)[keyof typeof PLANT_CODES]
type Plants = Plant[]
type Pots = Plants[]

export class Garden {
  constructor(private diagram: string, private students = DEFAULT_STUDENTS) {
    this.students.sort()
  }

  public plants(student: Student): Plants {
    const studentIndex = this.students.indexOf(student)
    const pots = this.diagram.split('\n').map((row) => row.split(''))
    const studentPots = pots.map((row) => row.slice(studentIndex * 2, studentIndex * 2 + 2))
    const studentPlants = studentPots.flatMap((pot) => pot.map((plant) => PLANT_CODES[plant as keyof typeof PLANT_CODES]))
    return studentPlants as Plants
  }
}
