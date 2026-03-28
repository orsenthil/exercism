export class GradeSchool {
  private db: Record<number, string[]> = {}
  studentPresent: string[] = []

  constructor() {
    this.db = {}
    this.studentPresent = []
  }

  roster() {
    return JSON.parse(JSON.stringify(this.db))
  }

  add(name: string, grade: number) {
      if (!this.db[grade]) {
          this.db[grade] = [name]
          this.studentPresent.push(name)
      } else {
          if (this.studentPresent.includes(name)) {
              // Remove all the students in the grade from this.db
              this.db = []
              return
          }
          this.db[grade].push(name)
          this.db[grade] = this.db[grade].sort()
          this.studentPresent.push(name)
      }
  }

  grade(grade: number) {
      let result: string[] = []
      for (let key in this.db) {
          if (parseInt(key) === grade) {
              result = this.db[key].sort()
          }
      }
      return JSON.parse(JSON.stringify(result))
  }
}
