<template>
  <div class="container">
    <div>
      <h1 class="title">natricons</h1>
      <button
        class="px-4 py-2 bg-primary text-white font-bold rounded-lg"
        @click="generateIframes()"
      >Randomize</button>
      <div class="flex flex-row flex-wrap mt-8">
        <iframe v-for="(frame, i) in iframes" :key="i" :src="frame.src" frameborder="0"></iframe>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  components: {},
  data() {
    return {
      iframes: []
    };
  },
  methods: {
    generateIframes() {
      const alphabet = "13456789abcdefghijkmnopqrstuwxyz";
      let randomAddress = "nano_";
      // After the ban prefix, we have 1 or 3, coinflip between them
      randomAddress += Math.floor(Math.random() * 2) == 0 ? "1" : "3";
      // Randomlys choose all other chars from the alphabet
      for (let i = 0; i < 59; i++) {
        const character = alphabet.charAt(Math.floor(Math.random() * 32));
        randomAddress += character;
      }
      let obj = {
        src: "http://localhost:8080/natricon?address=" + randomAddress
      };
      this.iframes.push(obj);
      console.log(this.iframes);
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
  font-weight: 300;
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
