<template>
  <div class="home container">
    <p>welcome to</p>
    <h1>Tiny Library</h1>
    <img
      alt="pile of books on a hand"
      class="giant-illustration container"
      src="../assets/book-background.jpg"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "Home"
});
</script>
<style scoped lang="scss">
.giant-illustration {
  position: absolute;
  bottom: 0;
  max-width: 100vw;
  left: 0;
  z-index: -100;
}
.home {
  min-height: 50vh;
  background-size: cover;
}
</style>
