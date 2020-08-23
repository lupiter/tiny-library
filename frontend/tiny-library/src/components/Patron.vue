<template>
  <div class="card">
    <div class="card-body">
      <h5 class="card-title">{{ patron.name }}</h5>
      <div class="card-text">{{ patron.card_number }}</div>
      <div class="card-text">
        Maximum loan period {{ patron.max_loan_days }} days
      </div>
      <div class="card-text">{{ patron.active ? "Active" : "Inactive" }}</div>
      <div class="catd-text">{{ patron.address }}</div>
      <button
        type="button"
        class="btn btn-primary"
        v-if="loans.length === 0"
        @click="fetchLoans"
      >
        Show loans
      </button>
      <button
        type="button"
        class="btn btn-primary"
        v-if="loans.length > 0"
        @click="removeLoans"
      >
        Close loans
      </button>
      <div class="container">
        <table class="table" v-if="loans.length > 0">
          <thead>
            <tr>
              <th scope="col">Book Title</th>
              <th scope="col">Book Author</th>
              <th scope="col">Out</th>
              <th scope="col">Expected</th>
              <th scope="col">Back</th>
              <th scope="col">Overdue</th>
            </tr>
          </thead>
          <tbody v-for="loan in loans" :key="loan.id">
            <Loan :identifier="loan.id" :showPatron="false" />
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { mapGetters, mapActions, mapMutations } from "vuex";
import { Patron, Loan as LoanModel } from "@/models/library";
import Loan from "@/components/Loan.vue"; // @ is an alias to /src

export default defineComponent({
  name: "Book",
  computed: {
    ...mapGetters(["patronById", "loansByPatronId"]),
    loans(): LoanModel[] {
      return this.loansByPatronId(this.identifier);
    },
    patron(): Patron | undefined {
      return this.patronById(this.identifier);
    }
  },
  props: {
    identifier: Number
  },
  components: {
    Loan
  },
  methods: {
    ...mapActions(["fetchLoansByPatron"]),
    ...mapMutations(["removeLoansByPatron"]),
    fetchLoans(): void {
      this.fetchLoansByPatron(this.identifier);
    },
    removeLoans(): void {
      this.removeLoansByPatron(this.identifier);
    }
  }
});
</script>

<style scoped lang="scss">
table {
  margin-top: 1rem;
}
</style>
