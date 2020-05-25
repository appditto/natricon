<template>
  <div class="w-full flex flex-col items-center">
    <hero />
    <how />
    <badges />
  </div>
</template>
<script>
import Hero from "~/components/Hero.vue";
import How from "~/components/How.vue";
import Badges from "~/components/Badges.vue";
export default {
  components: {
    Hero,
    How,
    Badges
  }
};
</script>