<template>
  <div class="home container">
    <div class="row">
      <p class="col">welcome to</p>
    </div>
    <div class="row">
      <h1 class="col">Tiny Library</h1>
    </div>
    <div class="row">
      <img
        alt="pile of books on a hand"
        class="giant-illustration container"
        src="../assets/book-background.jpg"
      />
    </div>
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
  p,
  h1 {
    margin: 2rem;
    text-align: center;
  }
}
</style>
