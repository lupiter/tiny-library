<template>
  <div class="home">
    <img alt="Tiny Library Logo" src="../assets/logo.png" />
    <h1>Welcome to Tiny Library</h1>


    <!-- <HelloWorld msg="Welcome to Your Vue.js + TypeScript App" /> -->
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
// import HelloWorld from "@/components/HelloWorld.vue"; // @ is an alias to /src

export default defineComponent({
  name: "Home",
  components: {
    // HelloWorld
  }
});
</script>
