<template>
  <div class="patrons container-fluid">
    <div class="row">
      <table class="table">
        <thead>
          <tr>
            <th scope="col">Name</th>
            <th scope="col">Card Number</th>
            <th scope="col">Lending Period</th>
            <th scope="col">Active</th>
            <th scope="col">Address</th>
            <th scope="col"></th>
          </tr>
        </thead>
        <tbody v-for="patron in patrons" :key="patron.id">
          <Patron :identifier="patron.id" />
        </tbody>
      </table>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Patron from "@/components/Patron.vue";
import { mapState, mapActions, mapMutations } from "vuex";

export default defineComponent({
  name: "Patrons",
  computed: mapState(["patrons"]),
  components: {
    Patron
  },
  methods: {
    ...mapActions(["fetchAllPatrons"]),
    ...mapMutations(["clearAllPatrons"])
  },
  mounted() {
    this.fetchAllPatrons();
  },
  unmounted() {
    this.clearAllPatrons();
  }
});
</script>
