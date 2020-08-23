<template>
  <tr>
    <th scope="row">{{ book.title }}</th>
    <td>{{ book.author }}</td>
    <td>{{ book.year }}</td>
    <td>{{ book.format }}</td>
    <td>{{ book.location }}</td>
  </tr>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { mapGetters } from "vuex";
import { Book } from "@/models/library";

export default defineComponent({
  name: "Book",
  computed: {
    ...mapGetters(["bookById"]),
    book(): Book | undefined {
      return this.bookById(this.identifier);
    }
  },
  props: {
    identifier: Number
  }
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss"></style>
