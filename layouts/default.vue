<template>
  <div>
    <navbar />
    <nuxt />
  </div>
</template>
<script>
import Navbar from "~/components/Navbar.vue";
export default {
  components: {
    Navbar
  }
};
</script>
<style>
html {
  font-family: "arboria", sans-serif;
}
</style>
