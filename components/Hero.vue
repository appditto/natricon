<template>
  <div class="w-full flex flex-col items-center pt-6 pb-8 md:pt-12 z-50" id="hero">
    <div class="n-container flex flex-col items-center px-5">
      <h1 class="text-4xl md:text-5xl text-center leading-none">
        meet your
        <br class="md:hidden" />
        <span class="relative inline-block">
          <span class="font-bold line-cyan">nano address</span>
        </span>
      </h1>
      <h2 class="text-xl md:text-2xl text-center mt-3">like you've never seen before.</h2>
      <button
        :class="isGeneratorOpen?'bg-cyan text-black btn-shadow-black':'bg-black text-white hover:text-cyan btn-shadow-cyan'"
        @click="isGeneratorOpen=!isGeneratorOpen"
        class="font-medium text-2xl rounded-full px-16 pt-1 pb-3 mt-5"
      >{{isGeneratorOpen?"close":"let's meet"}}</button>
    </div>
    <!-- Generator -->
    <div class="flex flex-row justify-center w-full h-0">
      <div class="generator-container z-40 mt-8 md:mt-10">
        <div
          :class="isGeneratorOpen ? 'scale-100 opacity-100 generator':'scale-0 opacity-0'"
          class="w-full h-full relative origin-top duration-300 justify-center items-center rounded-full bg-white mx-auto transform transition-all ease-out z-50 overflow-hidden"
        >
          <!-- Nano Address Form Group -->
          <label class="hidden" for="nanoAddressGroup">nano address</label>
          <form
            id="nanoAddressGroup"
            name="nanoAddressGroup"
            :class="generateInitiated?'scale-0':'scale-100'"
            class="w-full flex flex-col left-0 right-0 top-0 bottom-0 m-auto justify-center items-center px-5 md:px-6 lg:px-8 transform duration-200 ease-out absolute"
          >
            <label for="nanoAddressInput">nano address</label>
            <input
              :class="inputError?'border-red text-red':'border-black focus:bg-cyan'"
              class="w-full md:max-w-sm text-xl font-medium border-2 px-4 pt-1 pb-2 rounded-full my-1 transition-colors duration-200 ease-out"
              type="text"
              ref="nanoAddress"
              id="nanoAddressInput"
              name="nanoAddressInput"
              v-model="nanoAddress"
              placeholder="enter your address"
              @input="inputChange()"
            />
            <button
              @click.prevent="generateNatricon()"
              class="btn-2-shadow-cyan w-full md:max-w-sm btn text-xl font-medium hover:text-cyan bg-black text-white pt-1 pb-2 px-6 rounded-full mt-1"
            >meet!</button>
          </form>
          <!-- Randomize Button -->
          <button
            :class="generateInitiated?'scale-0':'scale-100 hover:scale-95'"
            @click.prevent="generateRandomNatricon()"
            class="btn-3-shadow-green bg-black left-half bottom-0 -translate-x-1/2 text-white hover:text-green btn-randomize absolute transform transition-all duration-200 text-lg md:text-xl mb-6 md:mb-8 lg:mb-10 pt-0_5 pb-1_5 md:pt-1 md:pb-2 px-4 md:px-5 lg:px-6 font-medium rounded-full"
          >randomize</button>
          <!-- Natricon Container -->
          <div
            v-if="generateInitiated"
            ref="natriconContainer"
            class="w-full h-full absolute top-0 left-0"
          ></div>
          <!-- Again Button -->
          <button
            :class="showAgainButton?'scale-100 hover:scale-95':'scale-0'"
            @click.prevent="resetProcess()"
            class="left-half bottom-0 -translate-x-1/2 btn-3-shadow-green hover:text-green bg-black text-white transform transition-all duration-200 md:text-xl mb-6 pt-0_5 pb-1_5 md:pt-1 md:pb-2 px-4 md:px-5 lg:px-6 font-medium border-black rounded-full absolute"
          >again!</button>
          <!-- Loading Animation -->
          <div
            v-if="generateInitiated"
            :class="natriconLoading?'scale-100 transition-all duration-200':'scale-0'"
            class="w-full h-full transform left-0 top-0 absolute"
          >
            <div
              class="absolute transform left-half top-half -translate-x-1/2 -translate-y-1/2 rounded-full bg-green green-circle"
            ></div>
            <div
              class="absolute transform left-half top-half -translate-x-1/2 -translate-y-1/2 rounded-full bg-brightPink brightPink-circle"
            ></div>
            <div
              class="absolute transform left-half top-half -translate-x-1/2 -translate-y-1/2 rounded-full bg-yellow yellow-circle"
            ></div>
            <div
              class="absolute transform left-half top-half -translate-x-1/2 -translate-y-1/2 rounded-full bg-cyan cyan-circle"
            ></div>
          </div>
          <!-- Received Animation -->
          <div
            v-if="generateInitiated"
            :class="receivedNatricon?'ray-margin-received':'ray-margin'"
            class="w-full h-full flex flex-row justify-center items-center left-0 top-0 absolute transition-all duration-1000 ease-out"
          >
            <div class="w-1/12 h-110 rounded-md bg-green"></div>
            <div class="w-1/12 h-120 rounded-md bg-brightPink"></div>
            <div class="w-1/12 h-130 rounded-md bg-yellow"></div>
            <div class="w-1/12 h-140 rounded-md bg-cyan"></div>
            <div class="w-1/12 h-110 rounded-md bg-green"></div>
            <div class="w-1/12 h-120 rounded-md bg-brightPink"></div>
            <div class="w-1/12 h-130 rounded-md bg-yellow"></div>
            <div class="w-1/12 h-140 rounded-md bg-cyan"></div>
            <div class="w-1/12 h-110 rounded-md bg-green"></div>
            <div class="w-1/12 h-120 rounded-md bg-brightPink"></div>
            <div class="w-1/12 h-130 rounded-md bg-yellow"></div>
            <div class="w-1/12 h-140 rounded-md bg-cyan"></div>
          </div>
        </div>
      </div>
    </div>
    <img
      class="w-full h-auto mt-12 md:hidden"
      :src="require('~/assets/images/illustrations/hero-mobile.svg')"
      alt="Hero Mobile"
    />
    <img
      class="w-full h-auto mt-12 hidden md:block lg:hidden"
      :src="require('~/assets/images/illustrations/hero-tablet.svg')"
      alt="Hero Tablet"
    />
    <img
      class="w-full h-auto mt-12 hidden lg:block"
      :src="require('~/assets/images/illustrations/hero-desktop.svg')"
      alt="Hero Desktop"
    />
  </div>
</template>
<script>
import { genAddress, validateAddress } from "~/plugins/address.js";
export default {
  data() {
    return {
      isGeneratorOpen: false,
      nanoAddress: "",
      inputError: false,
      generateInitiated: false,
      receivedNatricon: false,
      showAgainButton: false,
      natriconLoading: false,
      generatorClass: ""
    };
  },
  methods: {
    async generateNatricon() {
      let ref = this;
      if (validateAddress(ref.nanoAddress)) {
        ref.generateInitiated = true;
        setTimeout(() => {
          ref.natriconLoading = true;
        }, 100);
        const getNatriconResult = async () => {
          try {
            return await this.$axios.get(
              "https://natricon.com/api/v1/nano?svc=natricon.com&address=" +
                ref.nanoAddress
            );
          } catch (e) {
            console.error(e);
          }
        };
        const natriconResult = await getNatriconResult();
        if (natriconResult.data) {
          ref.receivedNatricon = true;
          setTimeout(() => {
            ref.$refs.natriconContainer.innerHTML = natriconResult.data;
            ref.natriconLoading = false;
            ref.showAgainButton = true;
          }, 300);
        } else {
          // Do something
        }
      } else {
        ref.inputError = true;
      }
    },
    generateRandomNatricon() {
      let randomAddress = genAddress();
      this.nanoAddress = randomAddress;
      this.generateNatricon();
    },
    inputChange() {
      if (this.inputError) {
        this.inputError = false;
      }
    },
    resetProcess() {
      this.showAgainButton = false;
      this.generateInitiated = false;
      this.receivedNatricon = false;
      this.$refs.natriconContainer.innerHTML = "";
    }
  }
};
</script>
<style scoped>
.btn {
  transition: all 0.2s ease-out;
  transform: scale(1);
}
.btn:hover {
  transform: scale(0.95);
}
.h-110 {
  height: 110%;
}
.h-120 {
  height: 120%;
}
.h-130 {
  height: 130%;
}
.h-140 {
  height: 140%;
}
.ray-margin-received {
  margin-top: 150%;
}
.ray-margin {
  margin-top: -150%;
}
.generator {
  box-shadow: 0rem 0.75rem 1.5rem 0rem rgba(0, 0, 0, 0.3),
    -0.9rem -0.9rem 0rem 0rem#66ffff, 0.65rem -0.7rem 0rem 0rem#FFA4F6,
    -0.5rem 0.85rem 0rem 0rem#FFEE52, 0.8rem 0.9rem 0rem 0rem#66FFB2;
  animation-name: shadow-animation;
  animation-duration: 6s;
  animation-iteration-count: infinite;
  animation-timing-function: ease-in-out;
}
@keyframes shadow-animation {
  0% {
    box-shadow: 0rem 0.5rem 1.5rem 0rem rgba(0, 0, 0, 0.3),
      -0.9rem -0.9rem 0rem 0rem#66ffff, 0.65rem -0.7rem 0rem 0rem#FFA4F6,
      -0.5rem 0.85rem 0rem 0rem#FFEE52, 0.8rem 0.9rem 0rem 0rem#66FFB2;
  }
  25% {
    box-shadow: 0rem 0.5rem 1.5rem 0rem rgba(0, 0, 0, 0.3),
      1.2rem 0.8rem 0rem 0rem#66ffff, -0.9rem 1rem 0rem 0rem#FFA4F6,
      0.7rem -1rem 0rem 0rem#FFEE52, -0.9rem -0.95rem 0rem 0rem#66FFB2;
  }
  50% {
    box-shadow: 0rem 0.5rem 1.5rem 0rem rgba(0, 0, 0, 0.3),
      -0.9rem -0.8rem 0rem 0rem#66ffff, 0.8rem -1.1rem 0rem 0rem#FFA4F6,
      -1.1rem 1rem 0rem 0rem#FFEE52, 1rem 0.9rem 0rem 0rem#66FFB2;
  }
  75% {
    box-shadow: 0rem 0.5rem 1.5rem 0rem rgba(0, 0, 0, 0.3),
      0.9rem 1.1rem 0rem 0rem#66ffff, -0.8rem 0.9rem 0rem 0rem#FFA4F6,
      0.9rem -0.9rem 0rem 0rem#FFEE52, -0.9rem -1.1rem 0rem 0rem#66FFB2;
  }
  100% {
    box-shadow: 0rem 0.5rem 1.5rem 0rem rgba(0, 0, 0, 0.3),
      -0.9rem -0.9rem 0rem 0rem#66ffff, 0.65rem -0.7rem 0rem 0rem#FFA4F6,
      -0.5rem 0.85rem 0rem 0rem#FFEE52, 0.8rem 0.9rem 0rem 0rem#66FFB2;
  }
}
.btn-shadow-cyan {
  box-shadow: -0.3rem 0.4rem 0rem 0rem#66ffff;
  transition: all 0.2s ease-out;
  transform: scale(1);
}
.btn-shadow-cyan:hover {
  box-shadow: 0rem 0rem 0rem 0rem#66ffff;
  transform: scale(0.95);
}
.btn-2-shadow-cyan {
  box-shadow: 0rem 0.4rem 0rem 0rem#66ffff;
}
.btn-2-shadow-cyan:hover {
  box-shadow: 0rem 0rem 0rem 0rem#66ffff;
}
.btn-3-shadow-green {
  box-shadow: 0rem 0.3rem 0rem 0rem#66FFB2;
}
.btn-3-shadow-green:hover {
  box-shadow: 0rem 0rem 0rem 0rem#66FFB2;
}
.btn-3-shadow-black {
  box-shadow: 0rem 0.3rem 0rem 0rem#000000;
}
.btn-3-shadow-black:hover {
  box-shadow: 0rem 0rem 0rem 0rem#000000;
}
.btn-shadow-black {
  box-shadow: -0.3rem 0.4rem 0rem 0rem#000000;
  transition: all 0.2s ease-out;
  transform: scale(1);
}
.btn-shadow-black:hover {
  box-shadow: 0rem 0rem 0rem 0rem#000000;
  transform: scale(0.95);
}
.line-cyan::after {
  display: block;
  position: absolute;
  width: calc(100% + 0.3rem);
  left: -0.15rem;
  content: "";
  height: 0.75rem;
  border-radius: 0.15rem;
  margin-left: auto;
  margin-right: auto;
  margin-top: -0.6rem;
  background-color: #66ffff;
  z-index: -1;
}
.bg-line {
  z-index: -1;
}
.generator-container {
  width: calc(100vw - 3rem);
  height: calc(100vw - 3rem);
  max-width: 25rem;
  max-height: 25rem;
}
@media only screen and (min-width: 768px) {
  .generator-container {
    width: 22vw;
    height: 22vw;
    min-width: 22rem;
    min-height: 22rem;
    max-width: 30vw;
    max-height: 30vw;
  }
}
.cyan-circle,
.brightPink-circle,
.green-circle,
.yellow-circle {
  width: 25%;
  height: 25%;
  margin-left: 0%;
  margin-top: 0%;
}
.cyan-circle {
  animation: cyan-animation;
  animation-duration: 2.2s;
  animation-iteration-count: infinite;
  animation-delay: -1s;
}
.green-circle {
  animation: cyan-animation;
  animation-duration: 2s;
  animation-iteration-count: infinite;
  animation-delay: -2s;
}
.brightPink-circle {
  animation: yellow-animation;
  animation-duration: 2.1s;
  animation-iteration-count: infinite;
  animation-delay: -2s;
}
.yellow-circle {
  animation: yellow-animation;
  animation-duration: 2.5s;
  animation-iteration-count: infinite;
  animation-delay: -1.2s;
}
@keyframes cyan-animation {
  0% {
    width: 25%;
    height: 25%;
    margin-left: 0%;
    margin-top: 0%;
  }
  25% {
    width: 32%;
    height: 32%;
    margin-left: -10%;
    margin-top: -10%;
  }
  50% {
    width: 23%;
    height: 23%;
    margin-left: 0%;
    margin-top: 0%;
  }
  75% {
    width: 30%;
    height: 30%;
    margin-left: 10%;
    margin-top: 10%;
  }
  100% {
    width: 25%;
    height: 25%;
    margin-left: 0%;
    margin-top: 0%;
  }
}
@keyframes yellow-animation {
  0% {
    width: 25%;
    height: 25%;
    margin-left: 0%;
    margin-top: 0%;
  }
  25% {
    width: 30%;
    height: 30%;
    margin-left: 10%;
    margin-top: -10%;
  }
  50% {
    width: 26%;
    height: 26%;
    margin-left: 0%;
    margin-top: 0%;
  }
  75% {
    width: 32%;
    height: 32%;
    margin-left: -10%;
    margin-top: 10%;
  }
  100% {
    width: 25%;
    height: 25%;
    margin-left: 0%;
    margin-top: 0%;
  }
}
.left-half {
  left: 50%;
}
.top-half {
  top: 50%;
}
</style>