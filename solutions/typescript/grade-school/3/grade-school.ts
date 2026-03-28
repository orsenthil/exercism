export class GradeSchool {
  private db: Record<number, string[]> = {}
  studentPresent: Record<string, number> = {}

  constructor() {
    this.db = {}
    this.studentPresent = {}
  }

  roster() {
    return JSON.parse(JSON.stringify(this.db))
  }

  add(name: string, grade: number) {
      if (this.studentPresent[name]) {
          this.db[this.studentPresent[name]] = []
          this.db[grade] = []
          return
      }

      if (!this.db[grade]) {
          this.db[grade] = [name]
          this.studentPresent[name] = grade
      } else {
          this.db[grade].push(name)
          this.db[grade] = this.db[grade].sort()
          this.studentPresent[name] = grade
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
