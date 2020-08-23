import { expect } from "chai";
import { shallowMount } from "@vue/test-utils";
import Book from "@/components/Book.vue";
import { Book as ModelBook } from "@/models/library";
import { createStore } from "vuex";

const identifier = 42;
const book = new ModelBook(
  "My Wonderful Life",
  "Juggles the Clown",
  "",
  "1991",
  "",
  [],
  identifier,
  0,
  "Bin",
  "Pamphlet"
);
const store = createStore({
  state() {
    return {
      books: [book]
    };
  },
  getters: {
    bookById: _ => (_: number): ModelBook | undefined => book
  }
});

describe("Book.vue", () => {
  it("renders book when passed id", () => {
    const wrapper = shallowMount(Book, {
      props: { identifier },
      global: {
        plugins: [store]
      }
    });
    expect(wrapper.text()).to.include(book.title);
    expect(wrapper.text()).to.include(book.author);
    expect(wrapper.text()).to.include(book.year);
    expect(wrapper.text()).to.include(book.format);
    expect(wrapper.text()).to.include(book.location);
  });
});
