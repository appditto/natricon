<template>
  <div class="container pb-24">
    <div>
      <h1 class="title text-primary text-5xl font-bold break-all">natricon</h1>
      <div class="flex flex-row flex-wrap justify-center">
        <button
          ref="btn2"
          class="px-4 py-2 m-2 bg-primary text-white text-xl font-bold rounded-lg transition-all duration-200 ease-out transform hover:scale-105"
          @click="generateRandomNatricon(); pulseIt2()"
        >Randomize 1</button>
        <button
          ref="btn"
          class="px-4 py-2 m-2 bg-primary text-white text-xl font-bold rounded-lg transition-all duration-200 ease-out transform hover:scale-105"
          @click="generateTenRandomNatricon(); pulseIt()"
        >Randomize 10</button>
      </div>
      <div v-if="natricons" class="flex flex-row justify-center flex-wrap">
        <natricon
          v-for="(natricon, i) in natricons"
          :key="i"
          :svg="natricon.svg"
          :address="natricon.address"
          :bodyH="natricon.bodyH"
          :bodyS="natricon.bodyS"
          :bodyV="natricon.bodyV"
          :hairH="natricon.hairH"
          :hairS="natricon.hairS"
          :hairV="natricon.hairV"
          :deltaH="natricon.deltaH"
          :deltaS="natricon.deltaS"
          :deltaV="natricon.deltaV"
          class="mx-2 my-4 px-4 pb-4"
        />
      </div>
    </div>
  </div>
</template>

<script>
import Natricon from "~/components/Natricon.vue";
export default {
  components: {
    Natricon
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
        .get("/api/random")
        .then(res => {
          this.natricons.unshift(res.data);
          console.log(this.natricons);
        })
        .catch(err => console.log(err));
      return;
    },
    pulseIt() {
      if (this.$refs.btn.classList.contains("pulse")) {
        this.$refs.btn.classList.remove("pulse");
        setTimeout(() => {
          this.$refs.btn.classList.add("pulse");
        }, 25);
      } else {
        this.$refs.btn.classList.add("pulse");
      }
    },
    pulseIt2() {
      if (this.$refs.btn2.classList.contains("pulse")) {
        this.$refs.btn2.classList.remove("pulse");
        setTimeout(() => {
          this.$refs.btn2.classList.add("pulse");
        }, 25);
      } else {
        this.$refs.btn2.classList.add("pulse");
      }
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

.links {
  padding-top: 15px;
}
button:hover,
button:focus {
  outline: none;
}
.pulse {
  animation: pulse-animation 0.5s;
}
@keyframes pulse-animation {
  0% {
    box-shadow: 0 0 0 0 rgba(136, 43, 255, 0.75);
  }
  100% {
    box-shadow: 0 0 0 1rem rgba(136, 43, 255, 0);
  }
}
</style>
