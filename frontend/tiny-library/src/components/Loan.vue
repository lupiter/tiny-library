<template>
  <tr>
    <th scope="row">{{ loan.book.title }}</th>
    <td>{{ loan.book.author }}</td>
    <td v-if="showPatron">{{ loan.patron.name }}</td>
    <td>{{ lent }}</td>
    <td>{{ due }}</td>
    <td>{{ returned }}</td>
    <td>{{ overdue }}</td>
  </tr>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { mapGetters } from "vuex";
import moment from "moment";
import { Loan } from "@/models/library";

export default defineComponent({
  name: "Loan",
  computed: {
    ...mapGetters(["loanById"]),
    loan(): Loan | undefined {
      return this.loanById(this.identifier);
    },
    overdue(): string {
      if (
        this.loan &&
        (this.loan.returned === undefined || this.loan.returned === "")
      ) {
        const due = moment(this.loan.due_back);
        const now = moment();
        if (due.isBefore(now)) {
          return due.toNow(true);
        }
      }
      return "";
    },
    lent(): string {
      if (this.loan) {
        return moment(this.loan.lent).format("D MMM YYYY");
      }
      return "";
    },
    due(): string {
      if (this.loan) {
        return moment(this.loan.due_back).format("D MMM YYYY");
      }
      return "";
    },
    returned(): string {
      if (this.loan && this.loan.returned) {
        return moment(this.loan.returned).format("D MMM YYYY");
      }
      return "";
    }
  },
  props: {
    identifier: Number,
    showPatron: Boolean
  }
});
</script>

<style scoped lang="scss"></style>
