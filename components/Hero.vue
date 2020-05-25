<template>
  <div class="w-full flex flex-col items-center pt-4 pb-8 md:pt-12">
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
        class="bg-black text-white btn-shadow-cyan font-medium text-2xl rounded-full px-16 pt-1 pb-3 mt-5"
      >let's meet</button>
    </div>
    <img
      class="w-full h-auto mt-12 md:hidden"
      :src="require('~/assets/images/illustrations/hero-mobile.svg')"
      alt="Hero Mobile"
    />
    <img
      class="w-full h-auto mt-12 hidden md:block"
      :src="require('~/assets/images/illustrations/hero-desktop.svg')"
      alt="Hero Desktop"
    />
  </div>
</template>
<style scoped>
.btn-shadow-cyan {
  box-shadow: -0.3rem 0.4rem 0rem 0rem#66ffff;
}
.line-cyan::before {
  display: block;
  position: absolute;
  width: calc(100% + 0.3rem);
  left: -0.15rem;
  content: "";
  height: 0.75rem;
  border-radius: 0.15rem;
  top: 100%;
  margin-left: auto;
  margin-right: auto;
  margin-top: -0.6rem;
  background-color: #66ffff;
  z-index: -1;
}
</style>