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

export default defineComponent({
  name: "Book",
  data() {
    return {
      book: {
        author: ""
      }
    };
  },
  props: {
    id: String
  }
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss"></style>
