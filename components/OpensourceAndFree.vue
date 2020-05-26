<template>
  <div class="w-full flex flex-col items-center py-8 md:py-24">
    <div class="n-container flex flex-col items-center px-5">
      <h3 class="text-4xl md:text-5xl text-center leading-tight">
        <span class="relative inline-block">
          <span class="font-bold line-cyan">open-source</span>
        </span>
        <br class="md:hidden" />&
        <span class="relative inline-block">
          <span class="font-bold line-lightPink">free</span>
        </span>
      </h3>
      <h4 class="text-xl md:text-2xl text-center mt-3">
        check out our code, see if you like it. contribute to it.
        it’s free.
        <br class="hidden md:block" />if you like it, you can donate to keep us going.
        <br class="hidden md:block" />donations of
        <b>2 nano</b> will get a
        <b>donor badge for 1 month</b>.
      </h4>
      <div class="w-full flex flex-row flex-wrap justify-center mt-2 relative">
        <button
          class="btn btn-shadow-cyan hover:text-cyan w-full md:w-56 bg-black text-white font-medium text-xl rounded-full px-6 pt-1 pb-3 mt-5 mx-3"
        >visit the repo</button>
        <button
          @click="isDropdownOpen?closeDonateDropdown():openDonateDropdown()"
          :class="isDropdownOpen?'bg-lightPink text-black btn-shadow-black':'bg-black text-white hover:text-lightPink btn-shadow-lightPink'"
          class="btn w-full md:w-56 font-medium text-xl rounded-full px-6 pt-1 pb-3 mt-5 mx-3"
        >{{isDropdownOpen?'close':'donate'}}</button>
      </div>
      <!-- Donate Dropdow Container-->
      <div class="flex flex-row justify-center bg-white relative">
        <!-- Donate Dropdown -->
        <div
          :class="isDropdownOpen?'h-72 border-black shadow-lightPink':'h-0 border-transparent'"
          class="absolute md:w-128 bg-white border-4 transition-all duration-300 ease-out overflow-hidden mt-5 rounded-lg"
        ></div>
      </div>
      <div class="w-full max-w-lg md:max-w-full px-4 md:px-24 md:hidden mt-12">
        <img
          class="w-full h-auto"
          :src="require('~/assets/images/illustrations/opensource-and-free-mobile.svg')"
          alt="Open-source and Free Mobile"
        />
      </div>
      <div class="w-full md:px-16 lg:px-24 mt-20 hidden md:block">
        <img
          class="w-full h-auto"
          :src="require('~/assets/images/illustrations/opensource-and-free-desktop.svg')"
          alt="Open-source and Free Desktop"
        />
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      isDropdownOpen: false
    };
  },
  methods: {
    openDonateDropdown() {
      this.isDropdownOpen = true;
    },
    closeDonateDropdown() {
      this.isDropdownOpen = false;
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
.btn-shadow-cyan {
  box-shadow: -0.3rem 0.4rem 0rem 0rem#66FFFF;
}
.btn-shadow-lightPink {
  box-shadow: -0.3rem 0.4rem 0rem 0rem#F199FF;
}
.shadow-lightPink {
  box-shadow: -0.5rem 0.6rem 0rem 0rem#F199FF;
}
.btn-shadow-black {
  box-shadow: -0.3rem 0.4rem 0rem 0rem#000000;
}
.btn-shadow-cyan:hover {
  box-shadow: 0rem 0rem 0rem 0rem#66FFFF;
}
.btn-shadow-lightPink:hover {
  box-shadow: 0rem 0rem 0rem 0rem#F199FF;
}
.btn-shadow-black:hover {
  box-shadow: 0rem 0rem 0rem 0rem#000000;
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
  margin-top: -0.9rem;
  background-color: #66ffff;
  z-index: -1;
}
.line-lightPink::after {
  display: block;
  position: absolute;
  width: calc(100% + 0.3rem);
  left: -0.15rem;
  content: "";
  height: 0.75rem;
  border-radius: 0.15rem;
  margin-left: auto;
  margin-right: auto;
  margin-top: -0.9rem;
  background-color: #f199ff;
  z-index: -1;
}
.border-transparent {
  border-color: rgba(0, 0, 0, 0);
}
</style>