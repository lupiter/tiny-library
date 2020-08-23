<template>
  <div id="app">
    <nav id="nav" class="navbar navbar-expand-sm navbar-light bg-light">
      <router-link to="/" class="navbar-brand brand">
        <img src="./assets/logo.png" width="30" height="30" class="d-inline-block align-top" alt="">
        Tiny Library
      </router-link>
      <button
        class="navbar-toggler"
        type="button"
        data-toggle="collapse"
        data-target="#navbarContents"
        aria-controls="navbarContents"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarContents">
        <div class="navbar-nav">
          <router-link to="/books" class="nav-link">Catalogue</router-link>
          <router-link to="/people" class="nav-link">Patrons</router-link>
          <router-link to="/loans" class="nav-link">Loans</router-link>
          <router-link to="/about" class="nav-link">About</router-link>
        </div>
      </div>
    </nav>
    <router-view />
  </div>
</template>

<style lang="scss">
#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  .brand {
    color: #4B40E0;
  }
  .giant-illustration {
    position: absolute;
    opacity: 0.5;
    bottom: 0;
    max-width: 100vw;
    left: 0;
    z-index: -100;
  }
  h1.brand {
    img {
      height: 3rem;
    }
  }
}
</style>
