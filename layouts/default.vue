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
      },
      // Mono Font
      {
        rel: "stylesheet",
        href:
          "https://fonts.googleapis.com/css2?family=Space+Mono:wght@400;700&display=swap"
      }
    ]
  }
};
</script>
<style>
html {
  font-family: "arboria", sans-serif;
}
pre,
code,
.space-mono {
  font-family: "Space Mono", monospace;
}
</style>
