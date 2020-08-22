<template>
  <div class="loans">
    <table class="table">
      <thead>
        <tr>
          <th scope="col">Book Title</th>
          <th scope="col">Book Author</th>
          <th scope="col">Patron</th>
          <th scope="col">Out</th>
          <th scope="col">Expected</th>
          <th scope="col">Back</th>
        </tr>
      </thead>
      <tbody v-for="loan in loans" :key="loan.id">
        <Loan id="loan.id" />
      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Loan from "@/components/Loan.vue"; // @ is an alias to /src
import { mapState, mapActions } from "vuex";

export default defineComponent({
  name: "Loans",
  computed: mapState(["loans"]),
  components: {
    Loan
  },
  methods: mapActions(["fetchAllLoans"]),
  mounted() {
    this.fetchAllLoans();
  }
});
</script>
