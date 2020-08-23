import { expect } from "chai";
import { shallowMount, mount } from "@vue/test-utils";
import { nextTick } from "vue";
import Patron from "@/components/Patron.vue";
import {
  Patron as ModelPatron,
  Loan as ModelLoan,
  Book as ModelBook
} from "@/models/library";
import { createStore } from "vuex";

const identifier = 42;
const book = new ModelBook(
  "My Wonderful Life",
  "Juggles the Clown",
  "",
  "1991",
  "Very Nice Books Inc",
  [],
  identifier,
  0,
  "Bin",
  "Pamphlet"
);
const patron = new ModelPatron(
  "Juggles the Clown",
  identifier,
  "123456789",
  "2000-01-01",
  "1 Princess Hwy, Sydney",
  true,
  8
);
const loan = new ModelLoan(
  identifier,
  book,
  patron,
  "1900-01-01T00:00:00.00+00:00",
  "1900-01-22T00:00:00.00+00:00",
  undefined
);
const store = createStore({
  state: {
    patrons: [patron],
    loans: [] as ModelLoan[]
  },
  getters: {
    patronById: () => (): ModelPatron | undefined => patron,
    loansByPatronId: state => (): ModelLoan[] => state.loans,
    loanById: () => (): ModelLoan | undefined => loan
  },
  mutations: {
    removeLoansByPatron(state) {
      state.loans = [];
    },
    addLoans(state, loans: ModelLoan[]) {
      state.loans = [...state.loans, ...loans];
    }
  },
  actions: {
    fetchLoansByPatron({ commit }) {
      commit("addLoans", [loan]);
    }
  }
});

describe("Patron.vue", () => {
  it("renders patron", () => {
    const wrapper = shallowMount(Patron, {
      props: { identifier },
      global: {
        plugins: [store]
      }
    });
    expect(wrapper.text()).to.include(patron.name);
    expect(wrapper.text()).to.not.include(patron.date_of_birth);
  });

  it("responds to load loans button", async () => {
    const wrapper = mount(Patron, {
      props: { identifier },
      global: {
        plugins: [store]
      }
    });
    expect(wrapper.text()).to.not.include(book.title);
    wrapper.find("button").trigger("click");
    await nextTick();
    expect(wrapper.text()).to.include(book.title);
  });
});
