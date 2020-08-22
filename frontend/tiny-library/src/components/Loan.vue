<template>
  <tr>
    <th scope="row">{{ loanById(id).book.title }}</th>
    <td>{{ loanById(id).book.author }}</td>
    <td>{{ loanById(id).patron.name }}</td>
    <td>{{ loanById(id).lent }}</td>
    <td>{{ loanById(id).dueBack }}</td>
    <td>{{ loanById(id).returned }}</td>
  </tr>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Loan } from "@/models/library";
import { mapGetters } from "vuex";

export default defineComponent({
  name: "Loan",
  computed: {
    ...mapGetters([
      'loanById'
      // ...
    ])
  },
  props: {
    id: Number
  }
});
</script>

<style scoped lang="scss"></style>
