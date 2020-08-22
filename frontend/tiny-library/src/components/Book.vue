<template>
  <tr>
    <th scope="row">{{ bookById(identifier).title }}</th>
    <td>{{ bookById(identifier).author }}</td>
    <td>{{ bookById(identifier).year }}</td>
    <td>{{ bookById(identifier).format }}</td>
    <td>{{ bookById(identifier).location }}</td>
  </tr>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { mapGetters } from "vuex";

export default defineComponent({
  name: "Book",
  computed: {
    ...mapGetters(["bookById"])
  },
  props: {
    identifier: Number
  }
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss"></style>
