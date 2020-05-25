<template>
  <div class="w-full flex flex-col items-center py-8 md:py-24">
    <div class="n-container flex flex-col items-center px-5">
      <h3 class="text-4xl md:text-5xl text-center leading-none">
        how does it
        <br class="md:hidden" />
        <span class="relative inline-block">
          <span class="font-bold line-redPink">work?</span>
        </span>
      </h3>
      <h4 class="text-xl md:text-2xl text-center mt-3">
        we take your nano address, apply some math,
        <br class="hidden md:block" />and it magically turns into
        <span class="font-bold">natricon</span>.
      </h4>
      <div class="w-full max-w-md md:max-w-full px-16 md:px-24 md:hidden mt-8">
        <img
          class="w-full h-auto"
          :src="require('~/assets/images/illustrations/natricon-formula-mobile.svg')"
          alt="Natricon Formula Mobile"
        />
      </div>
      <div class="w-full px-24 mt-12 hidden md:block">
        <img
          class="w-full h-auto"
          :src="require('~/assets/images/illustrations/natricon-formula-desktop.svg')"
          alt="Natricon Formula Desktop"
        />
      </div>
    </div>
  </div>
</template>
<style scoped>
.line-redPink::before {
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
  background-color: #fe98cd;
  z-index: -1;
}
</style>