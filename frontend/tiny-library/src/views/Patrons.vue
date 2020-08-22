<template>
  <div class="patrons">
    <table class="table">
      <thead>
        <tr>
          <th scope="col">Name</th>
          <th scope="col">Card Number</th>
          <th scope="col">Maximum Lending Period</th>
          <th scope="col">Active</th>
          <th scope="col">Address</th>
        </tr>
      </thead>
      <tbody v-for="patron in patrons" :key="patron.id">
        <Patron id="patron.id" />
      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Patron from "@/components/Patron.vue";
import { mapState, mapActions } from "vuex";

export default defineComponent({
  name: "Patrons",
  computed: mapState(["patrons"]),
  components: {
    Patron
  },
  methods: mapActions(["fetchAllPatrons"]),
  mounted() {
    this.fetchAllPatrons();
  }
});
</script>
