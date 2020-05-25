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
  },
  head: {
    link: [
      {
        rel: "stylesheet",
        href: "https://use.typekit.net/ypc1kli.css"
      }
    ]
  }
};
</script>
<style>
html {
  font-family: "arboria", sans-serif;
}
</style>
