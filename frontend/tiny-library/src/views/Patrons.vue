<template>
  <div class="patrons container-fluid">
    <div class="row row-cols-1 row-cols-md-2">
      <div v-for="patron in patrons" :key="patron.id" class="col mb-4">
        <Patron :identifier="patron.id" />
      </div>
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

<style scoped lang="scss">
.patrons {
  margin-top: 1rem;
}
</style>
