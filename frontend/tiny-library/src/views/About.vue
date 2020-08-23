<template>
  <div class="about container">
    <div class="row">
      <h1 class="col">Tiny Library</h1>
    </div>
    <p class="row">
      <a href="https://github.com/lupiter/tiny-library">Github</a>
    </p>
    <img
      alt="pile of books on a hand"
      class="giant-illustration container"
      src="../assets/book-background.jpg"
    />
  </div>
</template>

<style scoped lang="scss">
.giant-illustration {
  position: absolute;
  bottom: 0;
  max-width: 100vw;
  left: 0;
  z-index: -100;
}
.about {
  h1,
  p {
    margin: 1rem;
    text-align: center;
  }
}
</style>
