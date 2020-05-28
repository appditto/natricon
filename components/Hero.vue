<template>
  <div class="w-full flex flex-col items-center pt-6 pb-8 md:pt-12 z-50">
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
    <div class="w-full flex flex-row justify-center h-auto mt-12 hidden lg:block relative">
      <!-- Generator -->
      <div
        ref="generatorInside"
        :class="isGeneratorOpen?'scale-100 opacity-100 duration-500':'scale-0 opacity-0 duration-200'"
        class="flex flex-col justify-center items-center rounded-full bg-white shadow-xl mx-auto -mt-8 transform transition-all ease-out absolute generator z-50 overflow-hidden"
      >
        <!-- Nano Address Form Group -->
        <form
          :class="generateInitiated?'scale-0':'scale-100' "
          class="w-full flex flex-col justify-center px-12 transform duration-200 ease-out"
        >
          <input
            :class="inputError?'border-red text-red':'border-black'"
            class="w-full text-xl font-medium border-2 px-4 pt-1 pb-2 rounded-lg my-1 transition-colors duration-200 ease-out"
            type="text"
            ref="nanoAddress"
            id="nanoAddress"
            name="nanoAddress"
            v-model="nanoAddress"
            placeholder="enter your address"
            @input="inputChange()"
          />
          <button
            @click.prevent="generateNatricon()"
            class="w-full btn text-xl font-medium border-black hover:text-cyan hover:border-cyan border-2 bg-black text-white pt-1 pb-2 px-6 rounded-lg mt-1"
          >go!</button>
        </form>
        <div v-if="generateInitiated" ref="natriconContainer" class="w-full h-full absolute"></div>
        <!-- Received Animation -->
        <div
          v-if="generateInitiated"
          :class="receivedNatricon?'ray-margin-received':'ray-margin'"
          class="w-full h-full flex flex-row justify-center items-center left-0 top-0 absolute transition-all duration-1000 ease-out"
        >
          <div class="w-1/12 h-110 rounded-md bg-lime"></div>
          <div class="w-1/12 h-120 rounded-md bg-brightPink"></div>
          <div class="w-1/12 h-130 rounded-md bg-yellow"></div>
          <div class="w-1/12 h-140 rounded-md bg-cyan"></div>
          <div class="w-1/12 h-110 rounded-md bg-lime"></div>
          <div class="w-1/12 h-120 rounded-md bg-brightPink"></div>
          <div class="w-1/12 h-130 rounded-md bg-yellow"></div>
          <div class="w-1/12 h-140 rounded-md bg-cyan"></div>
          <div class="w-1/12 h-110 rounded-md bg-lime"></div>
          <div class="w-1/12 h-120 rounded-md bg-brightPink"></div>
          <div class="w-1/12 h-130 rounded-md bg-yellow"></div>
          <div class="w-1/12 h-140 rounded-md bg-cyan"></div>
        </div>
      </div>
      <!--  -->
      <img
        class="w-full h-auto"
        :src="require('~/assets/images/illustrations/hero-background-desktop.svg')"
        alt="Hero Background Desktop"
      />
      <img
        :class="isGeneratorOpen?'hero-left-desktop-open':'hero-left-desktop'"
        class="h-auto absolute top-0 left-0 transition-all duration-500 ease-out"
        :src="require('~/assets/images/illustrations/hero-left-desktop.svg')"
        alt="Hero Left Desktop"
        ref="heroLeftDesktop"
      />
      <img
        :class="isGeneratorOpen?'hero-right-desktop-open':'hero-right-desktop'"
        class="h-auto absolute top-0 right-0 transition-all duration-500 ease-out"
        :src="require('~/assets/images/illustrations/hero-right-desktop.svg')"
        alt="Hero Right Desktop"
        ref="heroRightDestop"
      />
    </div>
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
      receivedNatricon: false
    };
  },
  methods: {
    openGenerator() {
      this.isGeneratorOpen = true;
      this.$refs.heroLeftDesktop.style.marginRight = "0rem";
    },
    async generateNatricon() {
      let ref = this;
      if (validateAddress(ref.nanoAddress)) {
        ref.generateInitiated = true;
        const getNatriconResult = async () => {
          try {
            return await this.$axios.get(
              "https://natricon.com/api/v1/nano?address=" + ref.nanoAddress
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
          }, 300);
        } else {
          // Do something
        }
      } else {
        ref.inputError = true;
      }
    },
    inputChange() {
      if (this.inputError) {
        this.inputError = false;
      }
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
.generator {
  left: 50%;
  margin-left: calc(-10vw + 2rem);
  width: calc(20vw - 4rem);
  height: calc(20vw - 4rem);
}
.hero-left-desktop,
.hero-right-desktop {
  width: 55.5%;
}
.hero-left-desktop-open,
.hero-right-desktop-open {
  width: 40%;
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
.btn-shadow-cyan {
  box-shadow: -0.3rem 0.4rem 0rem 0rem#66ffff;
  transition: all 0.2s ease-out;
  transform: scale(1);
}
.btn-shadow-cyan:hover {
  box-shadow: 0rem 0rem 0rem 0rem#66ffff;
  transform: scale(0.95);
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
</style>