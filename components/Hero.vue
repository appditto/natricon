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
        class="flex flex-col justify-center items-center rounded-full bg-white shadow-xl mx-auto -mt-8 transform transition-all ease-out absolute generator z-50"
      >
        <!-- Nano Address Form Group -->
        <form class="w-full flex flex-col justify-center px-12">
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
      inputError: false
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
        ref.$refs.generatorInside.innerHTML = natriconResult.data;
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
  margin-left: -10vw;
  width: 20vw;
  height: 20vw;
}
.hero-left-desktop,
.hero-right-desktop {
  width: 55.5%;
}
.hero-left-desktop-open,
.hero-right-desktop-open {
  width: 40%;
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