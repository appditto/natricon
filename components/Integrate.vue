<template>
  <div class="w-full flex flex-col items-center py-8 md:py-24">
    <div class="n-container flex flex-col items-center px-5">
      <h3 class="text-4xl md:text-5xl text-center leading-none">
        wanna integrate
        <br class="md:hidden" />
        <span class="relative inline-block">
          <span class="font-bold line-yellow">natricon</span>
        </span>?
      </h3>
      <h4 class="text-xl md:text-2xl text-center mt-3">let’s get you started. it’ll be a breeze.</h4>
      <nuxt-link class="mx-3 mt-5" to="/documentation">
        <button
          class="bg-black text-white hover:text-yellow btn-shadow-yellow font-medium text-xl rounded-full px-12 pt-1 pb-3"
        >let's start</button>
      </nuxt-link>
      <div class="w-full max-w-md md:max-w-full px-8 md:px-24 md:hidden mt-8">
        <img
          class="w-full h-auto"
          :src="require('~/assets/images/illustrations/integrate-mobile.svg')"
          alt="Integrate Mobile"
        />
      </div>
      <div class="w-full md:px-8 lg:px-16 mt-16 hidden md:block">
        <img
          class="w-full h-auto"
          :src="require('~/assets/images/illustrations/integrate-desktop.svg')"
          alt="Integrate Desktop"
        />
      </div>
    </div>
  </div>
</template>
<style scoped>
.btn-shadow-yellow {
  box-shadow: -0.3rem 0.4rem 0rem 0rem#FFEE52;
  transition: all 0.2s ease-out;
  transform: scale(1);
}
.btn-shadow-yellow:hover {
  box-shadow: 0rem 0rem 0rem 0rem#FFEE52;
  transform: scale(0.95);
}
.line-yellow::after {
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
  background-color: #ffee52;
  z-index: -1;
}
</style>