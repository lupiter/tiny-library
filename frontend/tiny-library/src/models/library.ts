export class Book {
  constructor(
    public title: string,
    public author: string,
    public isbn: string,
    public year: string,
    public publisher: string,
    public tags: string[],
    public id: number,
    public max_loan_days: number,
    public location: string,
    public format: string
  ) {}
}

export class Patron {
  constructor(
    public name: string,
    public id: number,
    public card_number: string,
    public date_of_birth: string,
    public address: string,
    public active: boolean,
    public max_loan_days: number
  ) {}
}

export class Loan {
  constructor(
    public id: number,
    public book: Book,
    public patron: Patron,
    public lent: string,
    public due_back: string,
    public returned: string | undefined
  ) {}
}
