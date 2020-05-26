<template>
  <div>
    <navbar />
    <nuxt />
    <n-footer />
  </div>
</template>
<script>
import Navbar from "~/components/Navbar.vue";
import NFooter from "~/components/NFooter.vue";
export default {
  components: {
    Navbar,
    NFooter
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
