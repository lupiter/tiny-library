<template>
  <div class="books">
    <table class="table">
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
        <Book id="book.id"/>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Book from "@/components/Book.vue"; // @ is an alias to /src
import { mapState, mapActions } from "vuex";

export default defineComponent({
  name: "Books",
  computed: mapState([
    'books'
  ]),
  components: {
    Book
  },
  methods: mapActions([
      'fetchAllBooks'
  ])
});
</script>
