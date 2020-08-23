import { createStore } from "vuex";
import { Book, Patron, Loan } from "@/models/library";

// For debuggin with a local backend
const host = `http://localhost:8080`;
// For prod
// const host = `${window.location.protocol}//${window.location.hostname}`;

export default createStore({
  state: {
    books: [] as Book[],
    patrons: [] as Patron[],
    loans: [] as Loan[]
  },
  getters: {
    bookById: state => (id: number): Book | undefined => {
      return state.books.find((x: Book): boolean => x.id === id);
    },
    loanById: state => (id: number): Loan | undefined => {
      return state.loans.find((x: Loan): boolean => x.id === id);
    },
    patronById: state => (id: number): Patron | undefined => {
      return state.patrons.find((x: Patron): boolean => x.id === id);
    },
    loansByPatronId: state => (id: number): Loan[] => {
      return state.loans.filter((x: Loan): boolean => x.patron.id === id);
    }
  },
  mutations: {
    setBooks(state, books: Book[]) {
      state.books = books;
    },
    setPatrons(state, patrons: Patron[]) {
      state.patrons = patrons;
    },
    setLoans(state, loans: Loan[]) {
      state.loans = loans;
    },
    addLoans(state, loans: Loan[]) {
      state.loans = [...state.loans, ...loans];
    },
    clearAllBooks(state) {
      state.books = [];
    },
    clearAllPatrons(state) {
      state.patrons = [];
    },
    clearAllLoans(state) {
      state.loans = [];
    },
    removeLoansByPatron(state, patronId: number) {
      state.loans = state.loans.filter(
        (x: Loan): boolean => x.patron.id !== patronId
      );
    }
  },
  actions: {
    fetchAllBooks({ commit }) {
      fetch(`${host}/api/v0/books`)
        .then(response => response.json())
        .then(data => commit("setBooks", data));
    },
    fetchAllPatrons({ commit }) {
      fetch(`${host}/api/v0/people`)
        .then(response => response.json())
        .then(data => commit("setPatrons", data));
    },
    fetchAllLoans({ commit }) {
      fetch(`${host}/api/v0/loans`)
        .then(response => response.json())
        .then(data => commit("setLoans", data));
    },
    fetchLoansByPatron({ commit }, patronId: number) {
      fetch(`${host}/api/v0/people/${patronId}/loans`)
        .then(response => response.json())
        .then(data => commit("addLoans", data));
    }
  },
  modules: {}
});
