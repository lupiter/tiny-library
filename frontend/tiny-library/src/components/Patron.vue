<template>
  <tr>
    <th scope="row">{{ patronById(identifier).name }}</th>
    <td>{{ patronById(identifier).card_number }}</td>
    <td>{{ patronById(identifier).max_loan_days }}</td>
    <td>{{ patronById(identifier).active }}</td>
    <td>{{ patronById(identifier).address }}</td>
    <td>
      <button type="button" class="btn btn-primary" v-if="books.length == 0">
        Fetch loans
      </button>
      <!-- <div class="books" v-if="books.length > 0">
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
            <Book :identifier="book.id" />
          </tbody>
        </table>
      </div> -->
    </td>
  </tr>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { mapGetters, mapState } from "vuex";

export default defineComponent({
  name: "Book",
  computed: {
    ...mapGetters(["patronById"]),
    ...mapState(["books"])
  },
  props: {
    identifier: Number
  }
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss"></style>
