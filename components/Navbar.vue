<template>
  <div class="w-full flex flex-row justify-center items-center px-4 py-4 md:px-6 md:py-6">
    <div class="n-container flex flex-row justify-between items-center">
      <nuxt-link class="relative" to="/">
        <div class="w-36 md:w-40 h-auto line bga-cyan scaleY">
          <img
            class="w-full h-auto"
            :src="require('~/assets/images/logos/natricon-logo.svg')"
            alt="Natricon Logo"
          />
        </div>
      </nuxt-link>
      <div class="flex flex-row font-medium">
        <div class="relative">
          <a
            class="text-lg md:text-xl line-2 bga-redPink scaleY"
            href="https://github.com/appditto/natricon"
            target="_blank"
          >github</a>
        </div>
        <div class="ml-6 mr-1 relative">
          <nuxt-link to="/documentation">
            <span class="text-lg md:text-xl line-2 bga-redPink scaleY">api</span>
          </nuxt-link>
        </div>
      </div>
    </div>
  </div>
</template>
<style scoped>
.bga-cyan::after {
  background-color: #66ffff;
}
.bga-redPink::after {
  background-color: #fe98cd;
}
.line::after {
  display: block;
  position: absolute;
  width: calc(100% + 0.4rem);
  left: -0.2rem;
  content: "";
  height: 0.8rem;
  border-radius: 0.15rem;
  margin-left: auto;
  margin-right: auto;
  margin-top: -0.6rem;
  z-index: -1;
  transition: all 0.2s ease-out;
  transform-origin: center bottom;
  transform: scaleY(0);
}
.scaleY:hover::after {
  transform: scaleY(1);
}
.line-2::after {
  display: block;
  position: absolute;
  width: calc(100% + 0.4rem);
  left: -0.2rem;
  content: "";
  height: 0.6rem;
  border-radius: 0.15rem;
  margin-left: auto;
  margin-right: auto;
  margin-top: -0.6rem;
  z-index: -1;
  transition: all 0.2s ease-out;
  transform-origin: center bottom;
  transform: scaleY(0);
}
</style>