<template>
  <div class="home">
    <img alt="Tiny Library Logo" src="../assets/logo.png" />
    <h1>Welcome to Tiny Library</h1>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "Home"
});
</script>
<style scoped lang="scss">
.home {
  background-image: url("../assets/book-background.jpg");
}
</style>
