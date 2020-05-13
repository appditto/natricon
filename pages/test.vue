<template>
  <div class="pb-24 flex flex-row justify-center">
    <div>
      <div class="flex flex-row justify-center items-center mt-8">
        <h2
          class="title text-primary text-xl font-bold break-all pr-2"
        >Total: {{ calculatePossibleBodyCombinations()}}</h2>
        <h2 class="title text-primary text-xl font-bold break-all pr-2">Hue:</h2>
        <input
          placeholder="240"
          ref="hueInput"
          v-model="hue"
          type="number"
          maxlength="3"
          class="border-2 border-primary px-2 py-1 text-lg rounded-lg"
        />
      </div>
      <div v-for="(y, t) in 9" :key="t" class="flex flex-row justify-center flex-wrap mt-4">
        <natricon-test
          v-for="(x, i) in 7"
          :key="i"
          :bodyH="hue%360"
          :bodyS="(saturation - t*0.1)"
          :bodyV="(brightness - i*0.1)"
          class="mx-2 my-4 px-3 pb-4"
        />
      </div>
    </div>
  </div>
</template>

<script>
import NatriconTest from "~/components/NatriconTest.vue";
export default {
  components: {
    NatriconTest
  },
  data() {
    return {
      hue: 240,
      saturation: 1,
      brightness: 1
    };
  },
  methods: {
    changeHueValue() {
      var currentValue = this.$refs.hueInput.value;
      this.hue = currentValue;
      console.log(this.hue);
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
    },
    calculatePossibleBodyCombinations() {
      var minPB = 50;
      var maxPB = 240;
      var totalNumber = 0;
      for (var r = 0; r < 256; r++) {
        for (var g = 0; g < 256; g++) {
          totalNumber =
            totalNumber +
            Math.min(
              Math.sqrt(
                Math.max(
                  (maxPB * maxPB - 0.241 * r * r - 0.691 * g * g) / 0.068,
                  0
                )
              ),
              255
            ) -
            Math.max(
              Math.sqrt(
                Math.max(
                  (minPB * minPB - 0.241 * r * r - 0.691 * g * g) / 0.068,
                  0
                )
              ),
              0
            );
        }
      }
      return totalNumber.toFixed();
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
