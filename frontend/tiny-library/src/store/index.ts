import { createStore } from "vuex";
import { Book, Patron, Loan } from '@/models/library';

export default createStore({
  state: {
    books: [] as Book[],
    patrons: [] as Patron[],
    loans: [] as Loan[]
  },
  getters: {
    bookById: (state) => (id: number): Book | undefined => {
      return state.books.find((x: Book): boolean => x.id == id)
    },
    loanById: (state) => (id: number): Loan | undefined => {
      return state.loans.find((x: Loan): boolean => x.id == id)
    },
    patronById: (state) => (id: number): Patron | undefined => {
      return state.patrons.find((x: Patron): boolean => x.id == id)
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
    }
  },
  actions: {
    fetchAllBooks({ commit }) {
      fetch(`${window.location}/v0/books`)
        .then(response => response.json())
        .then(data => commit('setBooks', data));
    },
    fetchAllPatrons({ commit }) {
      fetch(`${window.location}/v0/people`)
        .then(response => response.json())
        .then(data => commit('setPatrons', data));
    },
    fetchAllLoans({ commit }) {
      fetch(`${window.location}/v0/loans`)
        .then(response => response.json())
        .then(data => commit('setLoans', data));
    },
    fetchLoansByPatron({ commit }, patronId: number) {
      fetch(`${window.location}/v0/people/${patronId}/loans`)
        .then(response => response.json())
        .then(data => commit('setLoans', data));
    }
  },
  modules: {}
});
