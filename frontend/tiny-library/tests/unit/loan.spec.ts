import { expect } from "chai";
import { shallowMount } from "@vue/test-utils";
import Loan from "@/components/Loan.vue";
import {
  Loan as ModelLoan,
  Book as ModelBook,
  Patron as ModelPatron
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
  state() {
    return {
      loans: [loan]
    };
  },
  getters: {
    loanById: () => (): ModelLoan | undefined => loan
  }
});

describe("Loan.vue", () => {
  it("renders loan", () => {
    const wrapper = shallowMount(Loan, {
      props: { identifier },
      global: {
        plugins: [store]
      }
    });
    expect(wrapper.text()).to.include(loan.book.title);
    expect(wrapper.text()).to.not.include(loan.book.location);
  });
});
