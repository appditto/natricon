<template>
  <div class="container">
    <div>
      <h1 class="title">natricon</h1>
      <div class="flex flex-row flex-wrap justify-center">
        <button
          class="px-4 py-2 mx-2 bg-primary text-white text-xl font-bold rounded-lg"
          @click="generateRandomNatricon()"
        >Randomize</button>
        <button
          class="px-4 py-2 mx-2 bg-primary text-white text-xl font-bold rounded-lg"
          @click="generateTenRandomNatricon()"
        >Randomize 10</button>
      </div>
      <div v-if="natricons" class="flex flex-row justify-center flex-wrap">
        <sample-natricon
          v-for="(natricon, i) in natricons.slice().reverse()"
          :key="i"
          :bodyColor="'#'+natricon.bodyColor"
          :hairColor="'#'+natricon.hairColor"
          class="w-56 h-56"
        />
      </div>
    </div>
  </div>
</template>

<script>
import SampleNatricon from "~/components/SampleNatricon.vue";
export default {
  components: {
    SampleNatricon
  },
  data() {
    return {
      natricons: []
    };
  },
  methods: {
    generateTenRandomNatricon() {
      for (let i = 0; i < 10; i++) {
        this.generateRandomNatricon();
      }
    },
    generateRandomNatricon() {
      this.$axios
        .get("http://localhost:8080/random")
        .then(res => {
          this.natricons.push(res.data);
        })
        .catch(err => console.log(err));
      return;
    }
  }
};
</script>

<style>
/* Sample `apply` at-rules with Tailwind CSS
.container {
  @apply min-h-screen flex justify-center items-center text-center mx-auto;
}
*/
.container {
  margin: 0 auto;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: start;
  text-align: center;
}

.title {
  font-family: "Quicksand", "Source Sans Pro", -apple-system, BlinkMacSystemFont,
    "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
  display: block;
  font-weight: 600;
  font-size: 100px;
  color: #35495e;
  letter-spacing: 1px;
}

.subtitle {
  font-weight: 300;
  font-size: 42px;
  color: #526488;
  word-spacing: 5px;
  padding-bottom: 15px;
}

.links {
  padding-top: 15px;
}
</style>
