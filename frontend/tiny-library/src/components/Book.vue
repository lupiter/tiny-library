<template>
  <tr>
    <th scope="row">{{ bookById(id).title }}</th>
    <td>{{ bookById(id).author }}</td>
    <td>{{ bookById(id).year }}</td>
    <td>{{ bookById(id).format }}</td>
    <td>{{ bookById(id).location }}</td>
  </tr>
</template>

<script lang="ts">
import { defineComponent } from "vue";
// import { Book } from "@/models/library";
import { mapGetters } from "vuex";

export default defineComponent({
  name: "Book",
  computed: {
    ...mapGetters([
      'bookById'
      // ...
    ])
  },
  props: {
    id: Number
  }
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss"></style>
