export class GradeSchool {
  db: Record<number, string[]> = {}


  constructor() {
    // prevent modifying the db object from outside
    this.db = {}
  }

  roster() {
    return this.db
  }

  add(name: string, grade: number) {
      if (!this.db[grade]) {
          this.db[grade] = [name]
      } else {
          this.db[grade].push(name)
          this.db[grade] = this.db[grade].sort()
      }
  }

  grade(grade: number) {
      let result: string[] = []
      for (let key in this.db) {
          if (parseInt(key) === grade) {
              result = this.db[key].sort()
          }
      }
      return result
  }
}
