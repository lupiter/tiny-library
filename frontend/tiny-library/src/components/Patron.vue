<template>
  <tr>
    <th scope="row">{{ patronById(id).name }}</th>
    <td>{{ patronById(id).cardNumber }}</td>
    <td>{{ patronById(id).maxLoanDays }}</td>
    <td>{{ patronById(id).format }}</td>
    <td>{{ patronById(id).active }}</td>
    <td>{{ patronById(id).address }}</td>
    <td>
      <button>Fetch loans</button>
      <div class="books" v-if="books.length > 0">
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
    </td>
  </tr>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Book, Patron } from "@/models/library";
import { mapGetters, mapState } from "vuex";

export default defineComponent({
  name: "Book",
  computed: {
    ...mapGetters([
      'patronById'
    ]),
    ...mapState([
      'books'
    ])
  },
  props: {
    id: Number
  }
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss"></style>
