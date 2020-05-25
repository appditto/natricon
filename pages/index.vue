<template>
  <div class="w-full flex flex-col items-center">
    <hero />
    <how />
    <badges />
    <available-on />
  </div>
</template>
<script>
import Hero from "~/components/Hero.vue";
import How from "~/components/How.vue";
import Badges from "~/components/Badges.vue";
import AvailableOn from "~/components/AvailableOn.vue";
export default {
  components: {
    Hero,
    How,
    Badges,
    AvailableOn
  }
};
</script>