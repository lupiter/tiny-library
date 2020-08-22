<template>
  <tr>
    <th scope="row">{{ loanById(identifier).book.title }}</th>
    <td>{{ loanById(identifier).book.author }}</td>
    <td>{{ loanById(identifier).patron.name }}</td>
    <td>{{ loanById(identifier).lent }}</td>
    <td>{{ loanById(identifier).dueBack }}</td>
    <td>{{ loanById(identifier).returned }}</td>
  </tr>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { mapGetters } from "vuex";

export default defineComponent({
  name: "Loan",
  computed: {
    ...mapGetters(["loanById"])
  },
  props: {
    identifier: Number
  }
});
</script>

<style scoped lang="scss"></style>
