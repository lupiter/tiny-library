
export class Book {
  constructor(
    public title: string, 
    public author: string, 
    public isbn: string, 
    public year: string, 
    public publisher: string, 
    public tags: string[], 
    public id: number, 
    public maxLoanDays: number, 
    public location: string, 
    public format: string) {
  }
}

export class Patron {
  constructor(
    public name: string,
    public id: number,
    public cardNumber: string,
    public dateOfBirth: string,
    public address: string,
    public active: boolean,
    public maxLoanDays: number
  ) {}
}

export class Loan {
  constructor(
    public id: number,
    public book: Book,
    public patron: Patron,
    public lent: string,
    public dueBack: string,
    public returned: string
  ){}
}