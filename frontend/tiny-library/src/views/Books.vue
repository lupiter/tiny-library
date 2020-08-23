<template>
  <div class="books container-fluid">
    <div class="row">
      <table class="table col">
        <thead>
          <tr>
            <th scope="col">Title</th>
            <th scope="col">Author</th>
            <th scope="col">Year</th>
            <th scope="col">Format</th>
            <th scope="col">Location</th>
          </tr>
        </thead>
        <tbody v-for="book in books" :key="book.id">
          <Book :identifier="book.id" />
        </tbody>
      </table>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Book from "@/components/Book.vue"; // @ is an alias to /src
import { mapState, mapActions, mapMutations } from "vuex";

export default defineComponent({
  name: "Books",
  computed: mapState(["books"]),
  components: {
    Book
  },
  methods: {
    ...mapActions(["fetchAllBooks"]),
    ...mapMutations(["clearAllBooks"])
  },
  mounted() {
    this.fetchAllBooks();
  },
  unmounted() {
    this.clearAllBooks();
  }
});
</script>
